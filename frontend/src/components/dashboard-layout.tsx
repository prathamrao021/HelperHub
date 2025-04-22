import { ReactNode } from "react"
import { Navbar } from "@/components/Navbar"
import { useAuth } from "@/contexts/auth-context"

export function DashboardLayout({ children }: { children: ReactNode }) {
    const { user, logout } = useAuth()

    let menuItems: { title: string; href: string }[] = []

    if (user?.userRole === "VOLUNTEER") {
        menuItems = [
            { title: "Dashboard", href: "/dashboard" },
            { title: "Opportunities", href: "/opportunities" },
            { title: "My Applications", href: "/applications" },
        ]
    } else if (user?.userRole === "ORGANIZATION_ADMIN") {
        menuItems = [
            { title: "Dashboard", href: "/dashboard" },
            { title: "Manage Projects", href: "/projects" },
        ]
    }
    return (
        <div className="min-h-screen bg-background">
            <Navbar
                menuItems={menuItems}
                showThemeToggle={true}
                showLogButton={true}
                buttonDisplay="Logout"
                onLogClick={() => logout()}
            />
            <main className="pt-16">
                {children}
            </main>
        </div>
    )
}