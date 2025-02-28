import { LoginForm } from "@/components/login-form"
import { Navbar } from "@/components/Navbar"

export default function Login() {
  return (
    <>
      <Navbar
        menuItems={[]}
        showThemeToggle={true}
        showLoginButton={true}
        onLoginClick={() => console.log("Login button clicked")}
      />
      <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
        <div className="w-full max-w-sm">
          <LoginForm />
        </div>
      </div>
    </>
  )
}