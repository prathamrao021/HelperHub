// src/contexts/auth-context.tsx
import { createContext, useContext, useState, useEffect, ReactNode } from "react"

// Define a more comprehensive user type that handles both roles
type User = {
  id: string
  email: string
  userRole: "VOLUNTEER" | "ORGANIZATION_ADMIN"
  // Shared properties
  fullName: string
  profilePicture?: string | null
  // Volunteer-specific properties
  bio?: string
  skills?: string[]
  location?: string
  weeklyHours?: number
  // Organization-specific properties
  organizationName?: string
  address?: string
  description?: string
  phoneNumber?: string
}

type AuthContextType = {
  user: User | null
  isAuthenticated: boolean
  isLoading: boolean
  login: (email: string, password: string, role?: string) => Promise<void>
  registerVolunteer: (volunteerData: any) => Promise<void>
  registerOrganization: (organizationData: any) => Promise<void>
  logout: () => void
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    // Check if user is stored in localStorage on mount
    const storedUser = localStorage.getItem("user")
    if (storedUser) {
      setUser(JSON.parse(storedUser))
    }
    setIsLoading(false)
  }, [])

  const login = async (email: string, password: string, role?: string) => {
    try {
      setIsLoading(true)
      // Replace with your API call - adjust the endpoint based on role if needed
      const response = await fetch("/api/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password, role }),
      })

      if (!response.ok) {
        throw new Error("Login failed")
      }

      const userData = await response.json()
      
      // Save to state and localStorage
      setUser(userData)
      localStorage.setItem("user", JSON.stringify(userData))
    } catch (error) {
      console.error("Login error:", error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const registerVolunteer = async (volunteerData: any) => {
    try {
      setIsLoading(true)
      const response = await fetch("/api/volunteers/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          ...volunteerData,
          userRole: "VOLUNTEER"
        }),
      })

      if (!response.ok) {
        throw new Error("Volunteer registration failed")
      }

      // Optionally auto-login after registration
      await login(volunteerData.email, volunteerData.password, "volunteer")
    } catch (error) {
      console.error("Registration error:", error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const registerOrganization = async (organizationData: any) => {
    try {
      setIsLoading(true)
      const response = await fetch("/api/organizations/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          ...organizationData,
          userRole: "ORGANIZATION_ADMIN"
        }),
      })

      if (!response.ok) {
        throw new Error("Organization registration failed")
      }

      // Optionally auto-login after registration
      await login(organizationData.email, organizationData.password, "organization")
    } catch (error) {
      console.error("Registration error:", error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const logout = () => {
    setUser(null)
    localStorage.removeItem("user")
    // Redirect to home page
    window.location.href = "/"
  }

  const value = {
    user,
    isAuthenticated: !!user,
    isLoading,
    login,
    registerVolunteer,
    registerOrganization,
    logout,
  }

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

export const useAuth = () => {
  const context = useContext(AuthContext)
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider")
  }
  return context
}