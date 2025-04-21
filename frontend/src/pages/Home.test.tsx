import { render, screen, fireEvent } from '@testing-library/react'
import { describe, it, expect, vi, beforeEach } from 'vitest'
import { MemoryRouter, useNavigate } from 'react-router-dom'
import { Home } from './Home'

// Mock useNavigate
vi.mock('react-router-dom', async () => {
  const actual = await vi.importActual('react-router-dom')
  return {
    ...actual,
    useNavigate: vi.fn(),
  }
})

describe('Home Component', () => {
  const mockedUseNavigate = vi.mocked(useNavigate)

  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('should render the Home component correctly', () => {
    render(
      <MemoryRouter>
        <Home />
      </MemoryRouter>
    )

    // Check for the Navbar and HeroSection render
    expect(screen.getByText('Register')).toBeInTheDocument()
    expect(screen.getByText('About')).toBeInTheDocument()
  })
})