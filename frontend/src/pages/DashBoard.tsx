import React from "react"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from "@/components/ui/tabs"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Navbar } from "@/components/Navbar"
import { useAuth } from "@/contexts/auth-context"
import { 
  Calendar, Clock, Users, Briefcase, 
  Building, Award
} from "lucide-react"

// Dummy chart component - replace with a real chart library like recharts
const Overview: React.FC<{ userRole: string }> = ({ userRole }) => {
  return (
    <div className="h-[200px] flex items-center justify-center bg-muted/20 rounded-md">
      <p className="text-muted-foreground">
        {userRole === "ORGANIZATION_ADMIN" 
          ? "Organization Activity Chart" 
          : "Volunteer Hours Chart"}
      </p>
    </div>
  )
}

type JobProps = {
  name: string;
  organization?: string;
  volunteer?: string;
  date: string;
  hours: number;
  avatarSrc?: string;
}

const JobItem: React.FC<JobProps> = ({ 
  name, organization, volunteer, date, hours, avatarSrc 
}) => {
  const displayName = organization || volunteer || "";
  
  return (
    <div className="flex items-center">
      <Avatar className="h-9 w-9">
        <AvatarImage src={avatarSrc} alt={displayName} />
        <AvatarFallback>{displayName[0]}</AvatarFallback>
      </Avatar>
      <div className="ml-4 space-y-1">
        <p className="text-sm font-medium leading-none">{name}</p>
        <p className="text-sm text-muted-foreground">{displayName}</p>
      </div>
      <div className="ml-auto text-right">
        <p className="text-sm font-medium">{date}</p>
        <p className="text-sm text-muted-foreground">{hours} hours</p>
      </div>
    </div>
  )
}

const RecentVolunteerJobs: React.FC = () => {
  return (
    <div className="space-y-4">
      <JobItem
        name="Environmental Cleanup"
        organization="Green Earth"
        date="Jan 23, 2025"
        hours={4}
      />
      <JobItem
        name="Food Bank Assistant"
        organization="Community Pantry"
        date="Jan 18, 2025"
        hours={6}
      />
      <JobItem
        name="Web Development"
        organization="Nonprofit Tech"
        date="Jan 15, 2025"
        hours={8}
      />
      <JobItem
        name="Youth Mentoring"
        organization="Future Leaders"
        date="Jan 10, 2025"
        hours={3}
      />
      <JobItem
        name="Animal Shelter Help"
        organization="Paws & Care"
        date="Jan 5, 2025"
        hours={5}
      />
    </div>
  )
}

const RecentOrganizationJobs: React.FC = () => {
  return (
    <div className="space-y-4">
      <JobItem
        name="Environmental Cleanup"
        volunteer="Sarah Johnson"
        date="Jan 23, 2025"
        hours={4}
      />
      <JobItem
        name="Food Distribution"
        volunteer="Michael Chen"
        date="Jan 21, 2025"
        hours={6}
      />
      <JobItem
        name="Website Development"
        volunteer="Alex Rodriguez"
        date="Jan 17, 2025"
        hours={12}
      />
      <JobItem
        name="Event Coordination"
        volunteer="Jamie Taylor"
        date="Jan 14, 2025"
        hours={8}
      />
      <JobItem
        name="Social Media Campaign"
        volunteer="Pat Wilson"
        date="Jan 10, 2025"
        hours={5}
      />
    </div>
  )
}

const VolunteerProfile: React.FC<{ user: any }> = ({ user }) => {
  return (
    <div className="space-y-6">
      <div className="flex flex-col items-center justify-center space-y-3 sm:flex-row sm:space-y-0 sm:space-x-4">
        <Avatar className="h-24 w-24">
          <AvatarImage src={user?.profilePicture} />
          <AvatarFallback>{user?.fullName?.[0] || "V"}</AvatarFallback>
        </Avatar>
        <div className="space-y-1 text-center sm:text-left">
          <h3 className="text-2xl font-bold">{user?.fullName || "Volunteer Name"}</h3>
          <p className="text-muted-foreground">{user?.email || "volunteer@example.com"}</p>
          <div className="flex items-center space-x-1">
            <Award className="h-4 w-4 text-primary" />
            <span className="text-sm font-medium">Top Contributor</span>
          </div>
        </div>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Personal Information</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">Full Name</p>
              <p>{user?.fullName || "John Doe"}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">Email</p>
              <p>{user?.email || "john.doe@example.com"}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">Phone</p>
              <p>{user?.phoneNumber || "(555) 123-4567"}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">Location</p>
              <p>{user?.location || "New York, NY"}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">Skills</p>
              <div className="flex flex-wrap gap-1">
                {(user?.skills || ["Web Development", "Teaching", "Event Planning"]).map(
                  (skill: string) => (
                    <span key={skill} className="inline-flex items-center rounded-full bg-primary/10 px-2.5 py-0.5 text-xs font-medium text-primary">
                      {skill}
                    </span>
                  )
                )}
              </div>
            </div>
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">Joined</p>
              <p>January 2025</p>
            </div>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Bio</CardTitle>
        </CardHeader>
        <CardContent>
          <p>{user?.bio || "Passionate about making a difference through volunteer work. Experienced in community outreach and event organization. Looking to connect with nonprofits that focus on education and environmental issues."}</p>
        </CardContent>
      </Card>
      
      <div className="flex justify-end">
        <Button>Edit Profile</Button>
      </div>
    </div>
  )
}

const OrganizationProfile: React.FC<{ user: any }> = ({ user }) => {
  return (
    <div className="space-y-6">
      <div className="flex flex-col items-center justify-center space-y-3 sm:flex-row sm:space-y-0 sm:space-x-4">
        <Avatar className="h-24 w-24">
          <AvatarImage src={user?.profilePicture} />
          <AvatarFallback>{user?.organizationName?.[0] || "O"}</AvatarFallback>
        </Avatar>
        <div className="space-y-1 text-center sm:text-left">
          <h3 className="text-2xl font-bold">{user?.organizationName || "Organization Name"}</h3>
          <p className="text-muted-foreground">{user?.email || "org@example.com"}</p>
          <div className="flex items-center space-x-1">
            <Building className="h-4 w-4 text-primary" />
            <span className="text-sm font-medium">Nonprofit Organization</span>
          </div>
        </div>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Organization Information</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">Organization Name</p>
              <p>{user?.organizationName || "Community Helpers"}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">Email</p>
              <p>{user?.email || "contact@communityhelpers.org"}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">Phone</p>
              <p>{user?.phoneNumber || "(555) 987-6543"}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">Address</p>
              <p>{user?.address || "123 Nonprofit Way, Charity City, ST 12345"}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">Admin Name</p>
              <p>{user?.fullName || "Jane Smith"}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">Founded</p>
              <p>2020</p>
            </div>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Organization Description</CardTitle>
        </CardHeader>
        <CardContent>
          <p>{user?.description || "Our organization is dedicated to improving local communities through various outreach programs, educational initiatives, and environmental projects. We work with volunteers to create meaningful impact and sustainable change."}</p>
        </CardContent>
      </Card>
      
      <div className="flex justify-end">
        <Button>Edit Profile</Button>
      </div>
    </div>
  )
}

export function Dashboard() {
  const { user } = useAuth();
  const isOrganization = user?.userRole === "ORGANIZATION_ADMIN";
  
  // Set default menu items based on user role
  const menuItems = isOrganization
    ? [
        { title: "Dashboard", href: "/dashboard" },
        { title: "Manage Projects", href: "/projects" },
        { title: "Find Volunteers", href: "/volunteers" },
      ]
    : [
        { title: "Dashboard", href: "/dashboard" },
        { title: "Find Opportunities", href: "/opportunities" },
        { title: "My Applications", href: "/applications" },
      ];
  
  return (
    <div className="min-h-screen bg-background mt-20">
      <Navbar
        menuItems={menuItems}
        showThemeToggle={true}
        showLoginButton={false}
      />
      
      <div className="container mx-auto px-4 py-8">
        <div className="flex flex-col space-y-6">
          <div className="flex items-center justify-between">
            <h1 className="text-3xl font-bold">
              {isOrganization ? "Organization Dashboard" : "Volunteer Dashboard"}
            </h1>
          </div>

          <Tabs defaultValue="overview" className="space-y-6">
            <TabsList>
              <TabsTrigger value="overview">Overview</TabsTrigger>
              <TabsTrigger value="profile">Profile</TabsTrigger>
            </TabsList>
            
            <TabsContent value="overview" className="space-y-6">
              {/* Stat Cards - Different for each role */}
              <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
                {isOrganization ? (
                  // Organization-specific stat cards
                  <>
                    <Card>
                      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle className="text-sm font-medium">
                          Active Projects
                        </CardTitle>
                        <Briefcase className="h-4 w-4 text-muted-foreground" />
                      </CardHeader>
                      <CardContent>
                        <div className="text-2xl font-bold">12</div>
                        <p className="text-xs text-muted-foreground">
                          +2 from last month
                        </p>
                      </CardContent>
                    </Card>
                    
                    <Card>
                      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle className="text-sm font-medium">
                          Registered Volunteers
                        </CardTitle>
                        <Users className="h-4 w-4 text-muted-foreground" />
                      </CardHeader>
                      <CardContent>
                        <div className="text-2xl font-bold">48</div>
                        <p className="text-xs text-muted-foreground">
                          +8 from last month
                        </p>
                      </CardContent>
                    </Card>
                    
                    <Card>
                      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle className="text-sm font-medium">
                          Total Hours Contributed
                        </CardTitle>
                        <Clock className="h-4 w-4 text-muted-foreground" />
                      </CardHeader>
                      <CardContent>
                        <div className="text-2xl font-bold">356</div>
                        <p className="text-xs text-muted-foreground">
                          +42 from last month
                        </p>
                      </CardContent>
                    </Card>
                    
                    <Card>
                      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle className="text-sm font-medium">
                          Upcoming Events
                        </CardTitle>
                        <Calendar className="h-4 w-4 text-muted-foreground" />
                      </CardHeader>
                      <CardContent>
                        <div className="text-2xl font-bold">3</div>
                        <p className="text-xs text-muted-foreground">
                          Next event in 5 days
                        </p>
                      </CardContent>
                    </Card>
                  </>
                ) : (
                  // Volunteer-specific stat cards
                  <>
                    <Card>
                      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle className="text-sm font-medium">
                          Total Hours
                        </CardTitle>
                        <Clock className="h-4 w-4 text-muted-foreground" />
                      </CardHeader>
                      <CardContent>
                        <div className="text-2xl font-bold">126</div>
                        <p className="text-xs text-muted-foreground">
                          +12 from last month
                        </p>
                      </CardContent>
                    </Card>
                    
                    <Card>
                      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle className="text-sm font-medium">
                          Projects Completed
                        </CardTitle>
                        <Briefcase className="h-4 w-4 text-muted-foreground" />
                      </CardHeader>
                      <CardContent>
                        <div className="text-2xl font-bold">24</div>
                        <p className="text-xs text-muted-foreground">
                          +3 from last month
                        </p>
                      </CardContent>
                    </Card>
                    
                    <Card>
                      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle className="text-sm font-medium">
                          Organizations
                        </CardTitle>
                        <Building className="h-4 w-4 text-muted-foreground" />
                      </CardHeader>
                      <CardContent>
                        <div className="text-2xl font-bold">8</div>
                        <p className="text-xs text-muted-foreground">
                          +2 from last month
                        </p>
                      </CardContent>
                    </Card>
                    
                    <Card>
                      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                        <CardTitle className="text-sm font-medium">
                          Upcoming Events
                        </CardTitle>
                        <Calendar className="h-4 w-4 text-muted-foreground" />
                      </CardHeader>
                      <CardContent>
                        <div className="text-2xl font-bold">4</div>
                        <p className="text-xs text-muted-foreground">
                          2 this week
                        </p>
                      </CardContent>
                    </Card>
                  </>
                )}
              </div>

              {/* Charts and Recent Jobs - Different for each role */}
              <div className="grid gap-4 md:grid-cols-1 lg:grid-cols-7">
                <Card className="col-span-4">
                  <CardHeader>
                    <CardTitle>
                      {isOrganization ? "Project Activity" : "Volunteer Activity"}
                    </CardTitle>
                  </CardHeader>
                  <CardContent>
                    <Overview userRole={user?.userRole || "VOLUNTEER"} />
                  </CardContent>
                </Card>
                
                <Card className="col-span-3">
                  <CardHeader>
                    <CardTitle>
                      {isOrganization ? "Recent Project Activity" : "Recent Volunteer Jobs"}
                    </CardTitle>
                    <CardDescription>
                      {isOrganization 
                        ? "Recent volunteer contributions to your projects"
                        : "You've completed 5 jobs this month"}
                    </CardDescription>
                  </CardHeader>
                  <CardContent>
                    {isOrganization ? <RecentOrganizationJobs /> : <RecentVolunteerJobs />}
                  </CardContent>
                </Card>
              </div>
            </TabsContent>
            
            <TabsContent value="profile">
              {isOrganization 
                ? <OrganizationProfile user={user} /> 
                : <VolunteerProfile user={user} />
              }
            </TabsContent>
          </Tabs>
        </div>
      </div>
    </div>
  )
}