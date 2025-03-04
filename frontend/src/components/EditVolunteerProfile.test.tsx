import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { EditVolunteerProfile } from "./EditVolunteerProfile";
import { useAuth } from "@/contexts/auth-context";
import api from "@/lib/axios";
import { MemoryRouter } from "react-router-dom";

// Mock API
vi.mock("@/lib/axios", () => ({
  default: {
    put: vi.fn(),
  },
}));

// Mock Auth Context
vi.mock("@/contexts/auth-context", () => ({
  useAuth: vi.fn(),
}));

describe("EditVolunteerProfile Component", () => {
  let mockUser: any;
  let mockUpdateUser: ReturnType<typeof vi.fn>;

  beforeEach(() => {
    vi.clearAllMocks();
    // Mock user data
    mockUser = {
      email: "test@volunteer.com",
      name: "Test Volunteer",
      phone: "1234567890",
      location: "Some City",
      bio_Data: "A passionate volunteer with 3 years of experience.",
      available_Hours: 10,
      category_List: ["Teaching", "Event Planning"],
    };

    mockUpdateUser = vi.fn();

    // Mock useAuth hook
    vi.mocked(useAuth).mockReturnValue({
        user: mockUser,
        isAuthenticated: true,
        isLoading: false,
        updateUser: mockUpdateUser,
        login: vi.fn(),           
        registerVolunteer: vi.fn(), 
        registerOrganization: vi.fn(), 
        logout: vi.fn()          
      });

    // Mock API success response
    vi.mocked(api.put).mockResolvedValue({ data: { success: true } });
  });

  /** Test 1: Ensure "Edit Profile" button is visible */
  it("renders the Edit Profile button", () => {
    render(
      <MemoryRouter>
        <EditVolunteerProfile />
      </MemoryRouter>
    );

    expect(screen.getByText("Edit Profile")).toBeInTheDocument();
  });

  /** Test 2: Clicking "Edit Profile" opens the dialog */
  it("opens the dialog when clicking Edit Profile", () => {
    render(
      <MemoryRouter>
        <EditVolunteerProfile />
      </MemoryRouter>
    );

    fireEvent.click(screen.getByText("Edit Profile"));
    expect(screen.getByText("Edit Volunteer Profile")).toBeInTheDocument();
  });

  /** Test 3: Ensure form fields are pre-filled with user data */
  it("pre-fills the form with user data", () => {
    render(
      <MemoryRouter>
        <EditVolunteerProfile />
      </MemoryRouter>
    );

    fireEvent.click(screen.getByText("Edit Profile"));

    expect(screen.getByDisplayValue(mockUser.name)).toBeInTheDocument();
    expect(screen.getByDisplayValue(mockUser.phone)).toBeInTheDocument();
    expect(screen.getByDisplayValue(mockUser.location)).toBeInTheDocument();
    expect(screen.getByDisplayValue(mockUser.bio_Data)).toBeInTheDocument();
    expect(screen.getByDisplayValue(mockUser.available_Hours.toString())).toBeInTheDocument();
  });

  /** Test 4: Allow editing form fields */
  it("allows editing form fields", () => {
    render(
      <MemoryRouter>
        <EditVolunteerProfile />
      </MemoryRouter>
    );

    fireEvent.click(screen.getByText("Edit Profile"));

    const nameInput = screen.getByDisplayValue(mockUser.name);
    fireEvent.change(nameInput, { target: { value: "Updated Volunteer" } });

    expect(nameInput).toHaveValue("Updated Volunteer");
  });
});
