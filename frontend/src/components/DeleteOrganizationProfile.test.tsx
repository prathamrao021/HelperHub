import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { DeleteOrganizationProfile } from "./DeleteOrganizationProfile";
import { useAuth } from "@/contexts/auth-context";
import api from "@/lib/axios";
import { MemoryRouter } from "react-router-dom";

// Mock API
vi.mock("@/lib/axios", () => ({
  default: {
    delete: vi.fn(),
  },
}));

// Mock Auth Context
vi.mock("@/contexts/auth-context", () => ({
  useAuth: vi.fn(),
}));


const mockNavigate = vi.fn();
vi.mock("react-router-dom", async () => {
  const actual = await vi.importActual("react-router-dom");
  return {
    ...actual,
    useNavigate: () => mockNavigate, 
  };
});

describe("DeleteOrganizationProfile Component", () => {
  let mockUser: any;
  let mockLogout: ReturnType<typeof vi.fn>;

  beforeEach(() => {
    vi.clearAllMocks(); 
    vi.spyOn(console, "error").mockImplementation(() => {}); 

    // Mock user data
    mockUser = {
      email: "test@org.com",
      name: "Test Organization",
    };

    mockLogout = vi.fn();

    // Mock useAuth hook
    vi.mocked(useAuth).mockReturnValue({
      user: mockUser,
      isAuthenticated: true,
      isLoading: false,                 
      login: vi.fn(),                   
      registerVolunteer: vi.fn(),       
      registerOrganization: vi.fn(),    
      updateUser: vi.fn(),              
      logout: mockLogout,
    });

    // Mock API success response
    vi.mocked(api.delete).mockResolvedValue({ data: { success: true } });
  });

  /** Test 1: Ensure "Delete Profile" button is visible */
  it("renders the Delete Profile button", () => {
    render(
      <MemoryRouter>
        <DeleteOrganizationProfile />
      </MemoryRouter>
    );

    expect(screen.getByText("Delete Profile")).toBeInTheDocument();
  });

  /** Test 2: Clicking "Delete Profile" opens the confirmation dialog */
  it("opens the confirmation dialog when clicking Delete Profile", () => {
    render(
      <MemoryRouter>
        <DeleteOrganizationProfile />
      </MemoryRouter>
    );

    fireEvent.click(screen.getByText("Delete Profile"));

    expect(screen.getByText("Are you Sure????")).toBeInTheDocument();
  });

  /** Test 3: Clicking "Delete" calls API, logs out user, and redirects */
  it("calls API, logs out, and redirects on successful deletion", async () => {
    render(
      <MemoryRouter>
        <DeleteOrganizationProfile />
      </MemoryRouter>
    );

    fireEvent.click(screen.getByText("Delete Profile"));
    fireEvent.click(screen.getByText("Delete"));

    await waitFor(() => {
      expect(api.delete).toHaveBeenCalledTimes(1);
      expect(api.delete).toHaveBeenCalledWith(`organizations/delete/${mockUser.email}`);
      expect(mockLogout).toHaveBeenCalledTimes(1);
      expect(mockNavigate).toHaveBeenCalledWith("/"); 
    });
  });

  /** Test 4: Handles API failure gracefully */
  it("handles API failure correctly", async () => {
    vi.mocked(api.delete).mockRejectedValue(new Error("API error"));

    render(
      <MemoryRouter>
        <DeleteOrganizationProfile />
      </MemoryRouter>
    );

    fireEvent.click(screen.getByText("Delete Profile"));
    fireEvent.click(screen.getByText("Delete"));

    await waitFor(() => {
      expect(api.delete).toHaveBeenCalledTimes(1);
      expect(api.delete).toHaveBeenCalledWith(`organizations/delete/${mockUser.email}`);
      expect(mockLogout).not.toHaveBeenCalled(); 
      expect(mockNavigate).not.toHaveBeenCalled(); 
    });
  });
});
