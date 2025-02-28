import { Navbar } from "@/components/Navbar"

const navigationItems = [
    {
        title: "Register",
        subItems: [
            {
                title: "Register as Organization",
                href: "/register/organization",
                description: "Manage projects, events, and recruit volunteers."
            },
            {
                title: "Register as Volunteer",
                href: "/register/volunteer",
                description: "Explore opportunities and track your volunteer hours."
            }
        ]
    },
    {
        title: "About",
        href: "/about"
    }
]

const Home = () => {
    return (
        <div className="h-screen w-full">
            <Navbar
                menuItems={navigationItems}
                showThemeToggle={true}
                showLoginButton={true}
                onLoginClick={() => console.log("Login button clicked")}
            />
            <div>Home</div>
        </div>
    )
}
export default Home