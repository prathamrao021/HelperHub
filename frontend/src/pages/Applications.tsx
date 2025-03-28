import { useState, useEffect } from "react"
import { useAuth } from "@/contexts/auth-context"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import { Trash2, Search, X, Calendar, FileText, Clock } from "lucide-react"
import { format } from "date-fns"
import { toast, Toaster } from "sonner"
import { DashboardLayout } from "@/components/dashboard-layout"
import api from "@/lib/axios"
import { Skeleton } from "@/components/ui/skeleton"
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog"

// Mock data for applications
const mockApplications = [
  {
    id: 1,
    volunteer_id: 101,
    opportunity_id: 201,
    opportunity_title: "Community Garden Clean-up",
    organization_name: "Helping Hands Foundation",
    status: "pending",
    cover_letter: "I'm excited about this opportunity because I have experience in gardening and want to help the community.",
    created_at: "2023-06-01T10:30:00Z",
    updated_at: "2023-06-01T10:30:00Z"
  },
  {
    id: 2,
    volunteer_id: 101,
    opportunity_id: 202,
    opportunity_title: "Dog Walking Volunteers",
    organization_name: "Animal Shelter Network",
    status: "approved",
    cover_letter: "I love animals and have pet-sitting experience. I would be perfect for this role!",
    created_at: "2023-05-15T14:20:00Z",
    updated_at: "2023-05-20T09:15:00Z"
  },
  {
    id: 3,
    volunteer_id: 101,
    opportunity_id: 203,
    opportunity_title: "Weekend Math Tutoring",
    organization_name: "Education for All",
    status: "rejected",
    cover_letter: "I have a degree in mathematics and would love to share my knowledge with students.",
    created_at: "2023-05-10T11:45:00Z",
    updated_at: "2023-05-12T11:45:00Z"
  },
  {
    id: 4,
    volunteer_id: 101,
    opportunity_id: 204,
    opportunity_title: "Coding Workshop Leaders",
    organization_name: "Tech for Good",
    status: "pending",
    cover_letter: "I'm a software engineer with experience teaching coding to beginners. I'd love to help with your workshops.",
    created_at: "2023-05-25T08:30:00Z",
    updated_at: "2023-05-25T08:30:00Z"
  }
]

export function ApplicationsPage() {
  const { user } = useAuth()
  const [applications, setApplications] = useState(mockApplications)
  const [filteredApplications, setFilteredApplications] = useState(mockApplications)
  const [searchQuery, setSearchQuery] = useState("")
  const [isLoading, setIsLoading] = useState(true)

  // Fetch applications
  useEffect(() => {
    const fetchApplications = async () => {
      setIsLoading(true)
      try {
        // In a real implementation, replace this with actual API call
        // const response = await api.get(`/applications/volunteer/${user.id}`)
        // setApplications(response.data)
        // setFilteredApplications(response.data)
        
        // Using mock data for now
        setTimeout(() => {
          setApplications(mockApplications)
          setFilteredApplications(mockApplications)
          setIsLoading(false)
        }, 1000)
      } catch (error) {
        console.error('Error fetching applications:', error)
        toast.error('Failed to load applications')
        setIsLoading(false)
      }
    }

    fetchApplications()
  }, [user])

  // Handle search
  useEffect(() => {
    if (searchQuery.trim() === "") {
      setFilteredApplications(applications)
    } else {
      const lowerCaseQuery = searchQuery.toLowerCase()
      const filtered = applications.filter(
        app => 
          app.opportunity_title.toLowerCase().includes(lowerCaseQuery) ||
          app.organization_name.toLowerCase().includes(lowerCaseQuery) ||
          app.status.toLowerCase().includes(lowerCaseQuery)
      )
      setFilteredApplications(filtered)
    }
  }, [searchQuery, applications])

  // Handle withdrawal (delete application)
  const handleWithdraw = async (applicationId: number) => {
    try {
      // In a real implementation:
      // await api.delete(`/applications/${applicationId}`)
      
      // Update local state
      const updatedApplications = applications.filter(app => app.id !== applicationId)
      setApplications(updatedApplications)
      toast.success('Application withdrawn successfully')
    } catch (error) {
      console.error('Error withdrawing application:', error)
      toast.error('Failed to withdraw application')
    }
  }

  // Get status badge styling
  const getStatusBadge = (status: string) => {
    switch (status.toLowerCase()) {
      case 'approved':
        return <Badge className="bg-green-500 hover:bg-green-600">Approved</Badge>
      case 'rejected':
        return <Badge variant="destructive">Rejected</Badge>
      case 'pending':
      default:
        return <Badge variant="outline" className="text-amber-500 border-amber-500">Pending</Badge>
    }
  }

  return (
    <DashboardLayout>
      <div className="container mx-auto py-6 space-y-6">
        <Toaster position="top-right" />
        
        <div className="flex flex-col md:flex-row justify-between gap-4">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">My Applications</h1>
            <p className="text-muted-foreground">
              View and manage your volunteer applications
            </p>
          </div>
          
          <div className="relative w-full md:w-64">
            <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="Search applications..."
              className="pl-8"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
            />
            {searchQuery && (
              <Button
                variant="ghost"
                className="absolute right-0 top-0 h-full px-3"
                onClick={() => setSearchQuery("")}
              >
                <X className="h-4 w-4" />
              </Button>
            )}
          </div>
        </div>

        {isLoading ? (
          <div className="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
            {[1, 2, 3].map((i) => (
              <Card key={i} className="overflow-hidden">
                <CardHeader className="p-0">
                  <Skeleton className="h-32 w-full" />
                </CardHeader>
                <CardContent className="p-6 space-y-4">
                  <Skeleton className="h-6 w-3/4" />
                  <Skeleton className="h-4 w-full" />
                  <Skeleton className="h-4 w-full" />
                  <div className="flex justify-between items-center">
                    <Skeleton className="h-8 w-24" />
                    <Skeleton className="h-10 w-28" />
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        ) : filteredApplications.length > 0 ? (
          <div className="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
            {filteredApplications.map((application) => (
              <Card key={application.id} className="overflow-hidden">
                <CardHeader>
                  <CardTitle>{application.opportunity_title}</CardTitle>
                  <CardDescription>{application.organization_name}</CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  <div className="flex justify-between items-center">
                    <div className="flex items-center gap-2">
                      <Clock className="h-4 w-4 text-muted-foreground" />
                      <span className="text-sm text-muted-foreground">
                        Applied: {format(new Date(application.created_at), 'MMM d, yyyy')}
                      </span>
                    </div>
                    {getStatusBadge(application.status)}
                  </div>
                  
                  <div className="space-y-2">
                    <div className="flex items-center gap-2">
                      <FileText className="h-4 w-4 text-muted-foreground" />
                      <h4 className="text-sm font-medium">Cover Letter</h4>
                    </div>
                    <p className="text-sm text-muted-foreground line-clamp-3">
                      {application.cover_letter}
                    </p>
                  </div>
                </CardContent>
                <CardFooter className="flex justify-between">
                  <Button 
                    variant="outline" 
                    onClick={() => toast.info('View details feature coming soon')}
                  >
                    View Details
                  </Button>
                  
                  {application.status.toLowerCase() === 'pending' && (
                    <AlertDialog>
                      <AlertDialogTrigger asChild>
                        <Button variant="destructive">
                          <Trash2 className="mr-2 h-4 w-4" />
                          Withdraw
                        </Button>
                      </AlertDialogTrigger>
                      <AlertDialogContent>
                        <AlertDialogHeader>
                          <AlertDialogTitle>Are you sure?</AlertDialogTitle>
                          <AlertDialogDescription>
                            This will permanently withdraw your application for {application.opportunity_title}.
                            This action cannot be undone.
                          </AlertDialogDescription>
                        </AlertDialogHeader>
                        <AlertDialogFooter>
                          <AlertDialogCancel>Cancel</AlertDialogCancel>
                          <AlertDialogAction 
                            onClick={() => handleWithdraw(application.id)}
                            className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
                          >
                            Withdraw
                          </AlertDialogAction>
                        </AlertDialogFooter>
                      </AlertDialogContent>
                    </AlertDialog>
                  )}
                </CardFooter>
              </Card>
            ))}
          </div>
        ) : (
          <div className="flex flex-col items-center justify-center py-12 px-4">
            <FileText className="h-16 w-16 text-muted-foreground mb-4" />
            <h3 className="text-xl font-semibold mb-2">No applications found</h3>
            <p className="text-muted-foreground text-center mb-6">
              {searchQuery 
                ? "No applications match your search. Try different keywords."
                : "You haven't applied to any opportunities yet."}
            </p>
            {searchQuery && (
              <Button 
                variant="outline"
                onClick={() => setSearchQuery("")}
              >
                Clear search
              </Button>
            )}
          </div>
        )}
      </div>
    </DashboardLayout>
  )
}