import { render, screen } from '@testing-library/react'
import { describe, it, expect, vi, beforeEach, MockedFunction } from 'vitest'
import { MemoryRouter, NavigateFunction, useNavigate } from 'react-router-dom'
import Login from './Login'
import { AuthProvider } from '@/contexts/auth-context'  // Adjust the path as necessary

// Mock useNavigate
vi.mock('react-router-dom', async () => {
  const actual = await vi.importActual('react-router-dom')
  return {
    ...actual,
    useNavigate: vi.fn(),
  }
})

vi.mock('@/contexts/auth-context', async () => {
  const actual = await vi.importActual('@/contexts/auth-context')
  return {
    ...actual, // Preserve actual AuthProvider export
    useAuth: vi.fn().mockReturnValue({
      user: {
        id: 123,
        name: 'Test User',
        email: 'test@example.com',
        userRole: 'VOLUNTEER'
      },
      isAuthenticated: true,
      isLoading: false,
      login: vi.fn(),
      registerVolunteer: vi.fn(),
      registerOrganization: vi.fn(),
      logout: vi.fn(),
      updateUser: vi.fn(),
    }), // Mock useAuth
  }
})

describe('Login Component', () => {
  let mockedUseNavigate: MockedFunction<NavigateFunction>;

  beforeEach(() => {
    mockedUseNavigate = vi.mocked(useNavigate());
    vi.clearAllMocks();
  });

  it.only('should render the Login component correctly', () => {
    render(
      <MemoryRouter>
        <AuthProvider>
          <Login />
        </AuthProvider>
      </MemoryRouter>
    )

    // Check for the Navbar render
    expect(screen.getByText('HELPERHUB')).toBeInTheDocument()

    // Check for the LoginForm render
    expect(screen.getByText('Welcome Back')).toBeInTheDocument()
    expect(screen.getByText('Sign in to your account to continue')).toBeInTheDocument()
  })
})