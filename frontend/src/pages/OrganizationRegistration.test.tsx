import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { OrganizationRegistration } from "./OrganizationRegistration";
import { useAuth } from "@/contexts/auth-context";
import { MemoryRouter } from "react-router-dom";

// Mock Auth Context
vi.mock("@/contexts/auth-context", () => ({
  useAuth: vi.fn(),
}));

// Mock useNavigate
const mockNavigate = vi.fn();
vi.mock("react-router-dom", async () => {
  const actual = await vi.importActual("react-router-dom");
  return {
    ...actual,
    useNavigate: () => mockNavigate,
  };
});

describe("OrganizationRegistration Component", () => {
  let mockRegisterOrganization: ReturnType<typeof vi.fn>;

  beforeEach(() => {
    vi.clearAllMocks();

    mockRegisterOrganization = vi.fn().mockResolvedValue({});

    // Mock useAuth hook
    vi.mocked(useAuth).mockReturnValue({
      user: null,
      isAuthenticated: false,
      isLoading: false,
      login: vi.fn(),
      registerVolunteer: vi.fn(),
      registerOrganization: mockRegisterOrganization,
      updateUser: vi.fn(),
      logout: vi.fn(),
    });
  });

  // Test 1: Renders the form with fields and submit button
  it("renders the form with fields and submit button", () => {
    render(
      <MemoryRouter>
        <OrganizationRegistration />
      </MemoryRouter>
    );

    // Check if all form fields are rendered
    expect(screen.getByPlaceholderText("you@example.com")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("••••••••")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Your Organization Name")).toBeInTheDocument();
    expect(screen.getByText("Upload your organization logo")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("+1 (555) 000-0000")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("123 Main St, City, Country")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Tell us about your organization's mission, goals, and the impact you want to make...")).toBeInTheDocument();

    // Check if submit button is rendered
    expect(screen.getByText("Complete Profile")).toBeInTheDocument();
  });

  // Test 2: Validates form fields and shows error message on invalid input
  it("validates form fields and shows error message on invalid input", async () => {
    render(
      <MemoryRouter>
        <OrganizationRegistration />
      </MemoryRouter>
    );

    // Enter invalid data
    fireEvent.change(screen.getByPlaceholderText("you@example.com"), { target: { value: "invalidemail" } });
    fireEvent.change(screen.getByPlaceholderText("••••••••"), { target: { value: "short" } });
    fireEvent.change(screen.getByPlaceholderText("Your Organization Name"), { target: { value: "A" } });
    fireEvent.change(screen.getByPlaceholderText("+1 (555) 000-0000"), { target: { value: "1" } });
    // fireEvent.change(screen.getByPlaceholderText("123 Main St, City, Country"), { target: { value: "short" } });
    // fireEvent.change(screen.getByPlaceholderText("Tell us about your organization's mission, goals, and the impact you want to make..."), { target: { value: "short" } });

    // Trigger form validation
    fireEvent.click(screen.getByText("Complete Profile"));

    await waitFor(() => {
      expect(screen.getByText((content, element) => content.includes("Please enter a valid email address"))).toBeInTheDocument();
      expect(screen.getByText((content, element) => content.includes("Password must be at least 6 characters"))).toBeInTheDocument();
      expect(screen.getByText((content, element) => content.includes("Organization name must be at least 2 characters"))).toBeInTheDocument();
    //   expect(screen.getByText((content, element) => content.includes("Phone must be at least 10 characters"))).toBeInTheDocument();
    //   expect(screen.getByText((content, element) => content.includes("Location must be at least 10 characters"))).toBeInTheDocument();
    //   expect(screen.getByText((content, element) => content.includes("Description must be at least 10 characters"))).toBeInTheDocument();
    });
  });
  it("validates max length for organization name", async () => {
    render(
      <MemoryRouter>
        <OrganizationRegistration />
      </MemoryRouter>
    );

    // Enter an organization name that exceeds the max length
    const longOrganizationName = 'A'.repeat(101);  // 101 characters long
    fireEvent.change(screen.getByPlaceholderText("Your Organization Name"), { target: { value: longOrganizationName } });

    // Trigger form validation
    fireEvent.click(screen.getByText("Complete Profile"));

    await waitFor(() => {
      expect(screen.getByText("Organization name must not exceed 100 characters")).toBeInTheDocument();
    });
  });
  // Test 3: Submits form and calls registerOrganization on valid input
  it("submits form and calls registerOrganization on valid input", async () => {
    render(
      <MemoryRouter>
        <OrganizationRegistration />
      </MemoryRouter>
    );

    // Fill the form with valid inputs
    fireEvent.change(screen.getByPlaceholderText("you@example.com"), { target: { value: "test@org.com" } });
    fireEvent.change(screen.getByPlaceholderText("••••••••"), { target: { value: "password123" } });
    fireEvent.change(screen.getByPlaceholderText("Your Organization Name"), { target: { value: "My Organization" } });
    fireEvent.change(screen.getByPlaceholderText("+1 (555) 000-0000"), { target: { value: "1234567890" } });
    fireEvent.change(screen.getByPlaceholderText("123 Main St, City, Country"), { target: { value: "123 Main St, City, Country" } });
    fireEvent.change(screen.getByPlaceholderText("Tell us about your organization's mission, goals, and the impact you want to make..."), { target: { value: "This is a test organization description." } });

    // // Simulate file upload for Organization Logo
    // const file = new File(["(⌐□_□)"], "logo.png", { type: "image/png" });
    // const inputElement = screen.getByLabelText("Organization Logo");
    // fireEvent.change(inputElement, { target: { files: [file] } });

    // Submit the form
    fireEvent.click(screen.getByText("Complete Profile"));

    await waitFor(() => {
      expect(mockRegisterOrganization).toHaveBeenCalledTimes(1);
      expect(mockRegisterOrganization).toHaveBeenCalledWith({
        email: "test@org.com",
        password: "password123",
        name: "My Organization",
        phone: "1234567890",
        location: "123 Main St, City, Country",
        description: "This is a test organization description.",
        // profilePicture: file,
      });

      expect(mockNavigate).toHaveBeenCalledWith("/dashboard");
    });
  });
});