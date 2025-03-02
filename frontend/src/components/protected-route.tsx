import { ReactNode } from "react"
import { Navigate, useLocation } from "react-router-dom"
import { useAuth } from "@/contexts/auth-context"

// Basic protected route - requires authentication
export function ProtectedRoute({ children }: { children: ReactNode }) {
  const { isAuthenticated, isLoading } = useAuth()
  const location = useLocation()

  if (isLoading) {
    return <div>Loading...</div>
  }
  
  if (!isAuthenticated) {
    return <Navigate to="/login" state={{ from: location }} replace />
  }

  return <>{children}</>
}

// Role-specific protected route
export function RoleProtectedRoute({ 
  children, 
  allowedRoles 
}: { 
  children: ReactNode, 
  allowedRoles: ("VOLUNTEER" | "ORGANIZATION_ADMIN")[] 
}) {
  const { user, isAuthenticated, isLoading } = useAuth()
  const location = useLocation()

  if (isLoading) {
    return <div>Loading...</div>
  }

  if (!isAuthenticated) {
    return <Navigate to="/login" state={{ from: location }} replace />
  }

  if (!user || !allowedRoles.includes(user.userRole)) {
    return <Navigate to="/unauthorized" replace />
  }

  return <>{children}</>
}