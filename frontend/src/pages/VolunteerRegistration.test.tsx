import React from "react";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { VolunteerRegistration } from "./VoluteerRegistration";
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

describe("VolunteerRegistration Component", () => {
  let mockRegisterVolunteer: ReturnType<typeof vi.fn>;

  beforeEach(() => {
    vi.clearAllMocks();

    mockRegisterVolunteer = vi.fn().mockResolvedValue({});

    // Mock useAuth hook
    vi.mocked(useAuth).mockReturnValue({
      user: null,
      isAuthenticated: false,
      isLoading: false,
      login: vi.fn(),
      registerVolunteer: mockRegisterVolunteer,
      registerOrganization: vi.fn(),
      updateUser: vi.fn(),
      logout: vi.fn(),
    });
  });

  it("renders the form with fields and submit button", () => {
    render(
      <MemoryRouter>
        <VolunteerRegistration />
      </MemoryRouter>
    );

    // Check if all form fields are rendered
    expect(screen.getByPlaceholderText("John Doe")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("abcd@gmail.com")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("••••••••")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("+1 (555) 000-0000")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Tell us about yourself, your experience, and what motivates you to volunteer...")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("City, Country")).toBeInTheDocument();
    expect(screen.getByText("Select skills")).toBeInTheDocument();
    expect(screen.getByText("Complete Profile")).toBeInTheDocument();
  });

  it("validates form fields and shows error message on invalid input", async () => {
    render(
      <MemoryRouter>
        <VolunteerRegistration />
      </MemoryRouter>
    );

    // Trigger form validation
    fireEvent.click(screen.getByText("Complete Profile"));

    await waitFor(() => {
      expect(screen.getByText("Full name must be at least 2 characters")).toBeInTheDocument();
      expect(screen.getByText("Please enter a valid email address")).toBeInTheDocument();
      expect(screen.getByText("Password must be at least 6 characters")).toBeInTheDocument();
      expect(screen.getByText("Phone number must be at least 10 digits")).toBeInTheDocument();
      expect(screen.getByText("Bio must be at least 50 characters")).toBeInTheDocument();
      expect(screen.getByText("Location must be at least 3 characters")).toBeInTheDocument();
      expect(screen.getByText("Select at least one skill")).toBeInTheDocument();
    });
  });
});