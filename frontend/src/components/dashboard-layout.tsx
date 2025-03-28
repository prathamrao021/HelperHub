import { ReactNode } from "react"
import { Navbar } from "@/components/Navbar"
import { useAuth } from "@/contexts/auth-context"

export function DashboardLayout({ children }: { children: ReactNode }) {
  const { user } = useAuth()
  
  return (
    <div className="min-h-screen bg-background">
      <Navbar 
        menuItems={[
          { title: "Dashboard", href: "/dashboard" },
          { title: "Opportunities", href: "/opportunities" },
          { title: "My Applications", href: "/applications" },
        ]}
        showThemeToggle={true}
        showLogButton={true}
        buttonDisplay="Profile"
      />
      <main className="pt-16">
        {children}
      </main>
    </div>
  )
}