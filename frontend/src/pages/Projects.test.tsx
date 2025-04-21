import { render, screen } from '@testing-library/react'
import { describe, it, expect, vi, beforeEach } from 'vitest'
import { ManageProjectsPage } from './Projects'
import { useAuth } from '@/contexts/auth-context'
import api from '@/lib/axios'
import { MemoryRouter } from 'react-router-dom'

// Mock useAuth
vi.mock('@/contexts/auth-context', () => ({
  useAuth: vi.fn(),
}))

// Mock axios
vi.mock('@/lib/axios', () => ({
  get: vi.fn(),
  post: vi.fn(),
}))

describe('Projects Component', () => {
  beforeEach(() => {
    vi.clearAllMocks()
    vi.mocked(useAuth).mockReturnValue({
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
    })
  })

  it('should render the Projects component correctly', () => {
    render(
        <MemoryRouter>
            <ManageProjectsPage />
        </MemoryRouter>
    )
    
  })
})