import { Button } from "@/components/ui/button"
import { useNavigate } from "react-router-dom"

export function UnauthorizedPage() {
  const navigate = useNavigate()
  
  return (
    <div className="flex flex-col items-center justify-center min-h-screen p-4">
      <h1 className="text-2xl font-bold mb-4">Access Denied</h1>
      <p className="text-muted-foreground mb-8">
        You don't have permission to access this page.
      </p>
      <Button onClick={() => navigate("/")}>
        Go to Home
      </Button>
    </div>
  )
}