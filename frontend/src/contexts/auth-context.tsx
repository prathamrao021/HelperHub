// src/contexts/auth-context.tsx
import { createContext, useContext, useState, useEffect, ReactNode } from "react"
import api from "@/lib/axios"

// Define a more comprehensive user type that handles both roles
type User = {
  id: number
  email: string
  name: string
  userRole: "VOLUNTEER" | "ORGANIZATION_ADMIN"
  // Shared properties
  profilePicture?: string | null
  // Volunteer-specific properties
  bio_Data?: string
  skills?: string[]
  location?: string
  weeklyHours?: number
  // Organization-specific properties
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
      let userData
      
      if (role === "VOLUNTEER") {
        const response = await api.post("/login/volunteer", {
          email,
          password,
          role
        })
        // Create enhanced user data with role
        userData = {
          ...response.data,
          userRole: "VOLUNTEER"
        }
      } else {
        const response = await api.post("/login/organization", {
          email,
          password,
          role
        })
        // Create enhanced user data with role
        userData = {
          ...response.data,
          userRole: "ORGANIZATION_ADMIN"
        }
        console.log("Organization login data:", userData)
      }
      
      // Save to state and localStorage - now userData includes the userRole
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
      const response = api.post("/volunteers/create", {
        ...volunteerData,
        userRole: "VOLUNTEER"
      })

      if ((await response).status !== 200) {
        throw new Error("Volunteer registration failed")
      }


      await login(volunteerData.email, volunteerData.password, "VOLUNTEER")
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
      const response = api.post("/organizations/create", {
        ...
        organizationData,
        userRole: "ORGANIZATION_ADMIN"
      }
      )

      if ((await response).status !== 200) {
        throw new Error("Organization registration failed")
      }

      // Optionally auto-login after registration
      await login(organizationData.email, organizationData.password, "ORGANIZATION_ADMIN")
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