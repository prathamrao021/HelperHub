import { render, screen, waitFor } from '@testing-library/react'
import { describe, it, expect, vi, beforeEach, MockedFunction } from 'vitest'
import { MemoryRouter, useNavigate } from 'react-router-dom'
import { ProjectDetailPage } from './ProjectDetail'
import { AuthProvider } from '@/contexts/auth-context' // Adjust the import path as necessary

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

describe('ProjectDetailPage Component', () => {
  let mockedUseNavigate: MockedFunction<any>;

  beforeEach(() => {
    mockedUseNavigate = vi.mocked(useNavigate());
    vi.clearAllMocks();
  });

  it('should render loading state initially', () => {
    render(
      <MemoryRouter>
        <AuthProvider>
          <ProjectDetailPage />
        </AuthProvider>
      </MemoryRouter>
    )

  })

  it('should render project details correctly', async () => {
    render(
      <MemoryRouter>
        <AuthProvider>
          <ProjectDetailPage />
        </AuthProvider>
      </MemoryRouter>
    )
  })
})