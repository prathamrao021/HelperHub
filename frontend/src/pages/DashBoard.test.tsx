import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { describe, it, expect, vi, beforeEach, beforeAll } from "vitest";
import { Dashboard } from "./DashBoard";
import { useAuth } from "@/contexts/auth-context";
import { MemoryRouter } from "react-router-dom";

// Mock `window.matchMedia` to prevent errors in Navbar
beforeAll(() => {
  Object.defineProperty(window, "matchMedia", {
    writable: true,
    value: vi.fn().mockImplementation((query) => ({
      matches: false,
      media: query,
      onchange: null,
      addListener: vi.fn(),
      removeListener: vi.fn(),
      addEventListener: vi.fn(),
      removeEventListener: vi.fn(),
      dispatchEvent: vi.fn(),
    })),
  });
});

// Mock Components
vi.mock("@/components/OrganizationProfile", () => ({
  OrganizationProfile: () => <div data-testid="organization-profile" />,
}));

vi.mock("@/components/VolunteerProfile", () => ({
  VolunteerProfile: () => <div data-testid="volunteer-profile" />,
}));

// Mock Auth Context
vi.mock("@/contexts/auth-context", () => ({
  useAuth: vi.fn(),
}));

describe("Dashboard Component", () => {
  let mockUser: any;
  let mockLogout: ReturnType<typeof vi.fn>;

  beforeEach(() => {
    vi.clearAllMocks();

    mockLogout = vi.fn();

    // Default user as volunteer
    mockUser = {
      email: "test@user.com",
      name: "Test User",
      userRole: "VOLUNTEER",
    };

    vi.mocked(useAuth).mockReturnValue({
      user: mockUser,
      isAuthenticated: true,
      logout: mockLogout,
      isLoading: false,
      login: vi.fn(),
      registerVolunteer: vi.fn(),
      registerOrganization: vi.fn(),
      updateUser: vi.fn(),
    });
  });

  /** Test 1: Ensure "Volunteer Dashboard" renders for volunteers */
  it("renders the Volunteer Dashboard for volunteers", () => {
    render(
      <MemoryRouter>
        <Dashboard />
      </MemoryRouter>
    );

    expect(screen.getByText("Volunteer Dashboard")).toBeInTheDocument();
    expect(screen.getByText("Find Opportunities")).toBeInTheDocument();
    expect(screen.getByText("My Applications")).toBeInTheDocument();
    // expect(screen.getByTestId("volunteer-profile")).toBeInTheDocument();
  });

  /** Test 2: Ensure "Organization Dashboard" renders for organizations */
  it("renders the Organization Dashboard for organizations", () => {  
    mockUser = { ...mockUser, userRole: "ORGANIZATION_ADMIN" };
    
    //  Re-mock useAuth() to force Dashboard to re-render
    vi.mocked(useAuth).mockReturnValue({
      user: mockUser,
      isAuthenticated: true,
      logout: mockLogout,
      isLoading: false,
      login: vi.fn(),
      registerVolunteer: vi.fn(),
      registerOrganization: vi.fn(),
      updateUser: vi.fn(),
    });

    render(
      <MemoryRouter>
        <Dashboard />
      </MemoryRouter>
    );

    expect(screen.getByText("Organization Dashboard")).toBeInTheDocument();
    expect(screen.getByText("Manage Projects")).toBeInTheDocument();

    // Ensure `OrganizationProfile` is rendered
    // expect(screen.getByTestId("organization-profile")).toBeInTheDocument();
  });

  /** Test 3: Logout button calls logout function */
  it("calls logout function when clicking Logout button", async () => {
    render(
      <MemoryRouter>
        <Dashboard />
      </MemoryRouter>
    );

    fireEvent.click(screen.getByText("Logout"));

    await waitFor(() => {
      expect(mockLogout).toHaveBeenCalledTimes(1);
    });
  });
});
