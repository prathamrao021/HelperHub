import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { Home } from '@pages/Home'
import Login from '@pages/Login'
import { ThemeProvider } from '@components/theme-provider'
import { VolunteerRegistration } from '@pages/VoluteerRegistration'
import { OrganizationRegistration } from '@pages/OrganizationRegistration'
import { AuthProvider } from '@/contexts/auth-context'
import { Dashboard } from '@/pages/DashBoard'

import { ProtectedRoute, RoleProtectedRoute } from '@/components/protected-route'
import {UnauthorizedPage} from '@/pages/UnauthorizedPage'
import {OpportunitiesPage} from '@/pages/Opportunities'
import {ManageProjectsPage} from '@/pages/ManageProject'
import {ApplicationsPage} from '@/pages/Applications'

export function App() {
  return (
    <ThemeProvider defaultTheme='system' storageKey='vite-ui-theme'>
      <AuthProvider>
        <Router>
          <Routes>
            {/* Public routes */}
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register/volunteer" element={<VolunteerRegistration />} />
            <Route path="/register/organization" element={<OrganizationRegistration />} />
            <Route path="/unauthorized" element={<UnauthorizedPage />} />
            
            
            {/* Protected routes - any authenticated user */}
            <Route 
              path="/dashboard" 
              element={
                <ProtectedRoute>
                  <Dashboard />
                </ProtectedRoute>
              } 
            />
            
            {/* Volunteer specific routes */}
            <Route 
              path="/opportunities" 
              element={
                <RoleProtectedRoute allowedRoles={["VOLUNTEER"]}>
                  <OpportunitiesPage />
                </RoleProtectedRoute>
              } 
            />

            <Route 
              path="/applications" 
              element={
                <RoleProtectedRoute allowedRoles={["VOLUNTEER"]}>
                  <ApplicationsPage />
                </RoleProtectedRoute>
              }
            />
            
            {/* Organization specific routes */}
            <Route 
              path="/manage-projects" 
              element={
                <RoleProtectedRoute allowedRoles={["ORGANIZATION_ADMIN"]}>
                  <ManageProjectsPage />
                </RoleProtectedRoute>
              } 
            />
          </Routes>
        </Router>
      </AuthProvider>
    </ThemeProvider>
  )
}