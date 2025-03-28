import { useState, useEffect } from "react"
import { useAuth } from "@/contexts/auth-context"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import { MapPin, Clock, Search, X } from "lucide-react"
import { format } from "date-fns"
import { toast, Toaster } from "sonner"
import { DashboardLayout } from "@/components/dashboard-layout"

// Mock data for opportunities
const mockOpportunities = [
  {
    id: 1,
    organization_id: 101,
    organization_name: "Helping Hands Foundation",
    title: "Community Garden Clean-up",
    description: "Join us for a weekend clean-up of the community garden. Tasks include weeding, planting new flowers, and general maintenance.",
    location: "Central Park, New York",
    start_date: "2023-06-15",
    end_date: "2023-06-16",
    created_at: "2023-05-20T10:30:00Z",
    updated_at: "2023-05-20T10:30:00Z"
  },
  {
    id: 2,
    organization_id: 102,
    organization_name: "Animal Shelter Network",
    title: "Dog Walking Volunteers",
    description: "Help exercise our shelter dogs by taking them for walks. Training provided for new volunteers.",
    location: "Main Street Shelter, Boston",
    start_date: "2023-06-10",
    end_date: "2023-07-10",
    created_at: "2023-05-18T14:20:00Z",
    updated_at: "2023-05-19T09:15:00Z"
  },
  {
    id: 3,
    organization_id: 103,
    organization_name: "Education for All",
    title: "Weekend Math Tutoring",
    description: "Provide math tutoring to underprivileged middle school students. All materials provided.",
    location: "Lincoln High School, Chicago",
    start_date: "2023-06-03",
    end_date: "2023-07-29",
    created_at: "2023-05-15T11:45:00Z",
    updated_at: "2023-05-15T11:45:00Z"
  },
  {
    id: 4,
    organization_id: 104,
    organization_name: "Tech for Good",
    title: "Coding Workshop Leaders",
    description: "Lead introductory coding workshops for teens. Knowledge of HTML, CSS, and basic JavaScript required.",
    location: "Virtual/Online",
    start_date: "2023-06-20",
    end_date: "2023-06-25",
    created_at: "2023-05-22T08:30:00Z",
    updated_at: "2023-05-22T15:10:00Z"
  },
  {
    id: 5,
    organization_id: 105,
    organization_name: "Food Bank Alliance",
    title: "Food Drive Coordinator",
    description: "Coordinate weekend food collection drives at local supermarkets. Training session required before first shift.",
    location: "Multiple Locations, San Francisco",
    start_date: "2023-07-01",
    end_date: "2023-07-31",
    created_at: "2023-05-25T09:00:00Z",
    updated_at: "2023-05-26T14:20:00Z"
  }
]

type Opportunity = typeof mockOpportunities[0]

export function OpportunitiesPage() {
  const { user } = useAuth()
  const [opportunities, setOpportunities] = useState<Opportunity[]>([])
  const [loading, setLoading] = useState(true)
  const [searchQuery, setSearchQuery] = useState("")

  // Fetch opportunities (using mock data for now)
  useEffect(() => {
    // Simulate API call
    setTimeout(() => {
      setOpportunities(mockOpportunities)
      setLoading(false)
    }, 800)
  }, [])

  // Apply filters to opportunities (search only)
  const filteredOpportunities = opportunities.filter(opp => {
    // Text search in title and description
    return searchQuery === "" || 
      opp.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
      opp.description.toLowerCase().includes(searchQuery.toLowerCase()) ||
      opp.location.toLowerCase().includes(searchQuery.toLowerCase())
  })

  // Clear search
  const clearSearch = () => {
    setSearchQuery("")
  }

  // Accept an opportunity
  const acceptOpportunity = async (id: number) => {
    try {
      // Here you would make an API call to accept the opportunity
      // For now, we'll simply remove it from the list
      setOpportunities(opportunities.filter(opp => opp.id !== id))
      
      toast.success("Application Submitted", {
        description: "Your application has been submitted successfully.",
      })
    } catch (error) {
      toast.error("Error", {
        description: "There was an error submitting your application. Please try again.",
      })
    }
  }

  // Helper function for formatting dates
  const formatDate = (dateString: string) => {
    try {
      return format(new Date(dateString), "MMM d, yyyy")
    } catch (error) {
      return dateString
    }
  }

  return (
    <DashboardLayout>
      <Toaster position="top-right" expand={false} richColors />
      
      <div className="container mx-auto py-6 space-y-6">
        <div className="flex flex-col space-y-4">
          <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
            <h1 className="text-3xl font-bold tracking-tight">Volunteer Opportunities</h1>
          </div>
          
          {/* Always visible search bar */}
          <div className="relative">
            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="Search opportunities by title, description or location..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="pl-10 w-full"
            />
            {searchQuery && (
              <Button 
                variant="ghost" 
                size="sm" 
                className="absolute right-1 top-1/2 transform -translate-y-1/2 h-7 w-7 p-0" 
                onClick={clearSearch}
              >
                <X className="h-4 w-4" />
              </Button>
            )}
          </div>
        </div>

        {loading ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {[1, 2, 3].map(i => (
              <Card key={i} className="animate-pulse">
                <CardHeader className="bg-muted h-14 rounded-t-lg" />
                <CardContent className="space-y-2 pt-4">
                  <div className="h-4 bg-muted rounded w-3/4" />
                  <div className="h-20 bg-muted rounded" />
                  <div className="h-4 bg-muted rounded w-1/2" />
                </CardContent>
                <CardFooter className="bg-muted h-12 rounded-b-lg" />
              </Card>
            ))}
          </div>
        ) : filteredOpportunities.length > 0 ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {filteredOpportunities.map(opportunity => (
              <Card key={opportunity.id} className="overflow-hidden hover:border-primary transition-colors">
                <CardHeader className="pb-2">
                  <CardTitle className="line-clamp-1">{opportunity.title}</CardTitle>
                  <CardDescription className="line-clamp-1">
                    {opportunity.organization_name}
                  </CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  <p className="text-sm text-muted-foreground line-clamp-3">
                    {opportunity.description}
                  </p>
                  
                  <div className="flex items-center gap-2">
                    <MapPin className="h-4 w-4 text-muted-foreground" />
                    <span className="text-sm">{opportunity.location}</span>
                  </div>
                  
                  <div className="flex items-center gap-2">
                    <Clock className="h-4 w-4 text-muted-foreground" />
                    <span className="text-sm">
                      {formatDate(opportunity.start_date)} - {formatDate(opportunity.end_date)}
                    </span>
                  </div>
                  
                  <div className="flex flex-wrap gap-2">
                    <Badge variant="outline">
                      {Math.ceil((new Date(opportunity.end_date).getTime() - new Date(opportunity.start_date).getTime()) / (1000 * 60 * 60 * 24))} days
                    </Badge>
                  </div>
                </CardContent>
                <CardFooter>
                  <Button 
                    className="w-full"
                    onClick={() => acceptOpportunity(opportunity.id)}
                  >
                    Apply Now
                  </Button>
                </CardFooter>
              </Card>
            ))}
          </div>
        ) : (
          <div className="text-center py-10">
            <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-muted mb-4">
              <Search className="h-6 w-6 text-muted-foreground" />
            </div>
            <h3 className="text-lg font-medium">No opportunities found</h3>
            <p className="text-muted-foreground mt-2">
              Try adjusting your search or check back later for new opportunities.
            </p>
            {searchQuery && (
              <Button 
                variant="outline" 
                onClick={clearSearch}
                className="mt-4"
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