import { render, screen, fireEvent } from '@testing-library/react'
import { describe, it, expect, vi, beforeEach } from 'vitest'
import { MemoryRouter, useNavigate } from 'react-router-dom'
import { UnauthorizedPage } from './UnauthorizedPage'

// Mock useNavigate
vi.mock('react-router-dom', async () => {
  const actual = await vi.importActual('react-router-dom')
  return {
    ...actual,
    useNavigate: vi.fn(),
  }
})

describe('UnauthorizedPage Component', () => {
  let mockedUseNavigate: ReturnType<typeof vi.fn>;

  beforeEach(() => {
    mockedUseNavigate = vi.mocked(useNavigate);
  });

  it('should render the UnauthorizedPage correctly', () => {
    render(
      <MemoryRouter>
        <UnauthorizedPage />
      </MemoryRouter>
    )
    
    // Check for message
    expect(screen.getByText('Access Denied')).toBeInTheDocument()
    expect(screen.getByText("You don't have permission to access this page.")).toBeInTheDocument()
    expect(screen.getByText('Go to Home')).toBeInTheDocument()
  })
})