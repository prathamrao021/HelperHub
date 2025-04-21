import { useState, useEffect } from "react"
import { useParams, useNavigate } from "react-router-dom"
import { useAuth } from "@/contexts/auth-context"
import { DashboardLayout } from "@/components/dashboard-layout"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Skeleton } from "@/components/ui/skeleton"
import { toast, Toaster } from "sonner"
import { format } from "date-fns"
import { MapPin, Clock, Calendar, Users, User, CalendarClock, FileText, Check, X, ArrowLeft, FileCheck, FileX } from "lucide-react"
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
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"

// Mock data for a single opportunity
const mockOpportunity = {
    id: 1,
    organization_id: 101,
    title: "Community Garden Clean-up",
    description: "Join us for a weekend clean-up of the community garden. Tasks include weeding, planting new flowers, and general maintenance. We're looking for volunteers who are enthusiastic about environmental conservation and are willing to get their hands dirty! No prior experience is required, just a willingness to help and learn. We'll provide all the necessary tools and equipment, as well as refreshments throughout the day. This is a great opportunity to meet like-minded individuals and contribute to making our community greener and more beautiful.",
    location: "Central Park, New York",
    start_date: "2023-06-15",
    end_date: "2023-06-16",
    created_at: "2023-05-20T10:30:00Z",
    updated_at: "2023-05-20T10:30:00Z",
    total_applications: 12,
    pending_applications: 5,
    accepted_applications: 4,
    rejected_applications: 3
};

// Mock data for applications
const mockApplications = [
    {
        id: 101,
        volunteer_id: 201,
        volunteer_name: "Emma Wilson",
        volunteer_email: "emma.wilson@example.com",
        volunteer_avatar: "/avatars/emma.jpg",
        opportunity_id: 1,
        status: "pending",
        cover_letter: "I'm very excited about this opportunity as I've been passionate about gardening since childhood. I've maintained my own garden for years and would love to contribute my skills to help improve our community space. I'm available both days and can bring my own gardening gloves if needed.",
        created_at: "2023-05-22T14:30:00Z",
        updated_at: "2023-05-22T14:30:00Z"
    },
    {
        id: 102,
        volunteer_id: 202,
        volunteer_name: "Marcus Johnson",
        volunteer_email: "marcus.j@example.com",
        volunteer_avatar: "/avatars/marcus.jpg",
        opportunity_id: 1,
        status: "pending",
        cover_letter: "I work in landscaping professionally and would like to donate my weekend to help beautify our community garden. I have experience with all aspects of garden maintenance and can help guide others if needed.",
        created_at: "2023-05-23T10:15:00Z",
        updated_at: "2023-05-23T10:15:00Z"
    },
    {
        id: 103,
        volunteer_id: 203,
        volunteer_name: "Sophia Lee",
        volunteer_email: "sophialee@example.com",
        volunteer_avatar: "/avatars/sophia.jpg",
        opportunity_id: 1,
        status: "pending",
        cover_letter: "As an environmental science student, I'm always looking for hands-on experiences to contribute to community environmental efforts. I've participated in similar clean-ups before and am passionate about sustainable gardening practices.",
        created_at: "2023-05-24T16:45:00Z",
        updated_at: "2023-05-24T16:45:00Z"
    },
    {
        id: 104,
        volunteer_id: 204,
        volunteer_name: "James Miller",
        volunteer_email: "jamesmiller@example.com",
        volunteer_avatar: "/avatars/james.jpg",
        opportunity_id: 1,
        status: "pending",
        cover_letter: "I'm new to gardening but very eager to learn and help out. I'm a hard worker and quick learner. Looking forward to spending time outdoors and making a positive impact in our community.",
        created_at: "2023-05-25T11:20:00Z",
        updated_at: "2023-05-25T11:20:00Z"
    },
    {
        id: 105,
        volunteer_id: 205,
        volunteer_name: "Olivia Garcia",
        volunteer_email: "olivia.g@example.com",
        volunteer_avatar: "/avatars/olivia.jpg",
        opportunity_id: 1,
        status: "pending",
        cover_letter: "I've been looking for volunteer opportunities to get more involved in the community. I have basic gardening experience from maintaining my apartment balcony garden and am excited to contribute to a larger project.",
        created_at: "2023-05-26T09:10:00Z",
        updated_at: "2023-05-26T09:10:00Z"
    },
    {
        id: 106,
        volunteer_id: 206,
        volunteer_name: "Daniel Kim",
        volunteer_email: "daniel.kim@example.com",
        volunteer_avatar: "/avatars/daniel.jpg",
        opportunity_id: 1,
        status: "accepted",
        cover_letter: "I'm a retired botanist who would love to share my knowledge and expertise. I've worked with community gardens before and find it very rewarding to help create green spaces for everyone to enjoy.",
        created_at: "2023-05-21T13:40:00Z",
        updated_at: "2023-05-22T09:20:00Z"
    },
    {
        id: 107,
        volunteer_id: 207,
        volunteer_name: "Ava Martinez",
        volunteer_email: "ava.m@example.com",
        volunteer_avatar: "/avatars/ava.jpg",
        opportunity_id: 1,
        status: "accepted",
        cover_letter: "I run a local gardening club and would love to contribute to this initiative. Several of our members have also applied, and we're excited about the possibility of working together on this community project.",
        created_at: "2023-05-21T15:30:00Z",
        updated_at: "2023-05-22T09:25:00Z"
    },
    {
        id: 108,
        volunteer_id: 208,
        volunteer_name: "Noah Wilson",
        volunteer_email: "noah.w@example.com",
        volunteer_avatar: "/avatars/noah.jpg",
        opportunity_id: 1,
        status: "accepted",
        cover_letter: "I've been participating in community garden projects for years and would love to contribute to this one. I have experience with planting, weeding, and general maintenance.",
        created_at: "2023-05-21T16:15:00Z",
        updated_at: "2023-05-22T09:30:00Z"
    },
    {
        id: 109,
        volunteer_id: 209,
        volunteer_name: "Isabella Brown",
        volunteer_email: "isabella.b@example.com",
        volunteer_avatar: "/avatars/isabella.jpg",
        opportunity_id: 1,
        status: "accepted",
        cover_letter: "I'm a landscape architecture student looking to gain practical experience. I'm particularly interested in community-based projects and would value the opportunity to contribute my design knowledge while learning from others.",
        created_at: "2023-05-21T17:45:00Z",
        updated_at: "2023-05-22T09:35:00Z"
    },
    {
        id: 110,
        volunteer_id: 210,
        volunteer_name: "William Taylor",
        volunteer_email: "william.t@example.com",
        volunteer_avatar: "/avatars/william.jpg",
        opportunity_id: 1,
        status: "rejected",
        cover_letter: "I'm interested in helping out, but I can only commit to a few hours on Saturday morning. I hope this partial availability can still be useful for your project.",
        created_at: "2023-05-21T12:30:00Z",
        updated_at: "2023-05-22T10:15:00Z"
    },
    {
        id: 111,
        volunteer_id: 211,
        volunteer_name: "Charlotte Davis",
        volunteer_email: "charlotte.d@example.com",
        volunteer_avatar: "/avatars/charlotte.jpg",
        opportunity_id: 1,
        status: "rejected",
        cover_letter: "I'm eager to help with the garden cleanup, but I should mention that I have severe pollen allergies. I'm willing to take allergy medication, but wanted to be upfront about this limitation.",
        created_at: "2023-05-21T13:15:00Z",
        updated_at: "2023-05-22T10:20:00Z"
    },
    {
        id: 112,
        volunteer_id: 212,
        volunteer_name: "Benjamin Moore",
        volunteer_email: "benjamin.m@example.com",
        volunteer_avatar: "/avatars/benjamin.jpg",
        opportunity_id: 1,
        status: "rejected",
        cover_letter: "I'm interested in this opportunity, but I'm not sure if I'll be in town that weekend. I'm waiting to confirm some travel plans and would need to let you know closer to the date.",
        created_at: "2023-05-21T14:45:00Z",
        updated_at: "2023-05-22T10:25:00Z"
    }
];

export function ProjectDetailPage() {
    const { projectId } = useParams();
    const navigate = useNavigate();
    const { user } = useAuth();
    const [opportunity, setOpportunity] = useState<typeof mockOpportunity | null>(null);
    const [applications, setApplications] = useState<typeof mockApplications>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [selectedTab, setSelectedTab] = useState("pending");

    // Get applications filtered by status
    const pendingApplications = applications.filter(app => app.status === "pending");
    const acceptedApplications = applications.filter(app => app.status === "accepted");
    const rejectedApplications = applications.filter(app => app.status === "rejected");

    // Fetch opportunity and applications data
    useEffect(() => {
        const fetchData = async () => {
            setIsLoading(true);
            try {
                // In a real implementation, these would be API calls
                // const opportunityResponse = await api.get(`/opportunities/${projectId}`);
                // const applicationsResponse = await api.get(`/opportunities/${projectId}/applications`);

                // Using mock data for now
                setTimeout(() => {
                    setOpportunity(mockOpportunity);
                    setApplications(mockApplications);
                    setIsLoading(false);
                }, 1000);
            } catch (error) {
                console.error("Error fetching project data:", error);
                toast.error("Failed to load project data");
                setIsLoading(false);
            }
        };

        fetchData();
    }, [projectId]);

    // Handle accepting an application
    const handleAcceptApplication = async (applicationId: number) => {
        try {
            // In a real implementation, this would be an API call
            // await api.put(`/applications/${applicationId}`, { status: "accepted" });

            // Update local state
            setApplications(applications.map(app =>
                app.id === applicationId
                    ? { ...app, status: "accepted" }
                    : app
            ));

            toast.success("Application accepted successfully");
        } catch (error) {
            console.error("Error accepting application:", error);
            toast.error("Failed to accept application");
        }
    };

    // Handle rejecting an application
    const handleRejectApplication = async (applicationId: number) => {
        try {
            // In a real implementation, this would be an API call
            // await api.put(`/applications/${applicationId}`, { status: "rejected" });

            // Update local state
            setApplications(applications.map(app =>
                app.id === applicationId
                    ? { ...app, status: "rejected" }
                    : app
            ));

            toast.success("Application rejected successfully");
        } catch (error) {
            console.error("Error rejecting application:", error);
            toast.error("Failed to reject application");
        }
    };

    if (isLoading) {
        return (
            <DashboardLayout>
                <div className="container mx-auto py-6 space-y-8">
                    <Skeleton className="h-12 w-3/4 max-w-3xl" />
                    <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                        <div className="md:col-span-2 space-y-4">
                            <Skeleton className="h-40 w-full" />
                            <div className="space-y-2">
                                <Skeleton className="h-4 w-full" />
                                <Skeleton className="h-4 w-full" />
                                <Skeleton className="h-4 w-3/4" />
                            </div>
                        </div>
                        <div className="space-y-4">
                            <Skeleton className="h-32 w-full" />
                            <Skeleton className="h-32 w-full" />
                        </div>
                    </div>
                </div>
            </DashboardLayout>
        );
    }

    if (!opportunity) {
        return (
            <DashboardLayout>
                <div className="container mx-auto py-12 text-center">
                    <h2 className="text-2xl font-bold mb-4">Project Not Found</h2>
                    <p className="text-muted-foreground mb-6">The project you're looking for doesn't exist or you don't have permission to view it.</p>
                    <Button onClick={() => navigate('/projects')}>
                        <ArrowLeft className="mr-2 h-4 w-4" />
                        Back to Projects
                    </Button>
                </div>
            </DashboardLayout>
        );
    }

    return (
        <DashboardLayout>
            <Toaster position="top-right" richColors />

            <div className="container mx-auto py-6 space-y-8">
                {/* Header with back button and title */}
                <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
                    <div className="flex items-center gap-2">
                        <Button
                            variant="outline"
                            size="icon"
                            onClick={() => navigate('/projects')}
                            className="h-9 w-9"
                        >
                            <ArrowLeft className="h-4 w-4" />
                        </Button>
                        <h1 className="text-3xl font-bold tracking-tight">{opportunity.title}</h1>
                    </div>

                    <Badge className="bg-blue-500 hover:bg-blue-600">
                        {pendingApplications.length} Pending Application{pendingApplications.length !== 1 && 's'}
                    </Badge>
                </div>

                {/* Main content */}
                <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
                    {/* Project details section */}
                    <div className="lg:col-span-2 space-y-6">
                        <Card>
                            <CardHeader>
                                <CardTitle>Opportunity Details</CardTitle>
                                <CardDescription>Complete information about this volunteer opportunity</CardDescription>
                            </CardHeader>
                            <CardContent className="space-y-6">
                                <div className="space-y-4">
                                    <div className="flex items-start gap-2">
                                        <MapPin className="h-5 w-5 text-muted-foreground mt-0.5" />
                                        <div>
                                            <h3 className="font-semibold">Location</h3>
                                            <p className="text-muted-foreground">{opportunity.location}</p>
                                        </div>
                                    </div>

                                    <div className="flex items-start gap-2">
                                        <Calendar className="h-5 w-5 text-muted-foreground mt-0.5" />
                                        <div>
                                            <h3 className="font-semibold">Date</h3>
                                            <p className="text-muted-foreground">
                                                {format(new Date(opportunity.start_date), "MMMM d, yyyy")}
                                                {opportunity.start_date !== opportunity.end_date &&
                                                    ` - ${format(new Date(opportunity.end_date), "MMMM d, yyyy")}`
                                                }
                                            </p>
                                        </div>
                                    </div>
                                </div>

                                <div>
                                    <h3 className="font-semibold mb-2">Description</h3>
                                    <p className="text-muted-foreground">{opportunity.description}</p>
                                </div>
                            </CardContent>
                        </Card>

                        {/* Applications tabs */}
                        <Card>
                            <CardHeader>
                                <CardTitle>Applications</CardTitle>
                                <CardDescription>Manage volunteer applications for this opportunity</CardDescription>
                            </CardHeader>
                            <CardContent className="p-0">
                                <Tabs defaultValue={selectedTab} onValueChange={setSelectedTab}>
                                    <div className="px-6 pt-2">
                                        <TabsList className="grid grid-cols-3 w-full">
                                            <TabsTrigger value="pending" className="relative">
                                                Pending
                                                {pendingApplications.length > 0 && (
                                                    <Badge className="ml-2 bg-blue-500 hover:bg-blue-600">
                                                        {pendingApplications.length}
                                                    </Badge>
                                                )}
                                            </TabsTrigger>
                                            <TabsTrigger value="accepted">
                                                Accepted
                                                {acceptedApplications.length > 0 && (
                                                    <Badge className="ml-2 bg-green-500 hover:bg-green-600">
                                                        {acceptedApplications.length}
                                                    </Badge>
                                                )}
                                            </TabsTrigger>
                                            <TabsTrigger value="rejected">
                                                Rejected
                                                {rejectedApplications.length > 0 && (
                                                    <Badge className="ml-2 bg-red-500 hover:bg-red-600">
                                                        {rejectedApplications.length}
                                                    </Badge>
                                                )}
                                            </TabsTrigger>
                                        </TabsList>
                                    </div>

                                    <TabsContent value="pending" className="m-0">
                                        <div className="divide-y">
                                            {pendingApplications.length > 0 ? (
                                                pendingApplications.map((application) => (
                                                    <ApplicationCard
                                                        key={application.id}
                                                        application={application}
                                                        onAccept={handleAcceptApplication}
                                                        onReject={handleRejectApplication}
                                                        status="pending"
                                                    />
                                                ))
                                            ) : (
                                                <div className="py-12 text-center">
                                                    <div className="inline-flex h-12 w-12 items-center justify-center rounded-full bg-muted mb-4">
                                                        <User className="h-6 w-6 text-muted-foreground" />
                                                    </div>
                                                    <h3 className="text-lg font-medium">No pending applications</h3>
                                                    <p className="text-muted-foreground mt-1 mb-4 max-w-md mx-auto">
                                                        There are no pending applications for this opportunity at the moment.
                                                    </p>
                                                </div>
                                            )}
                                        </div>
                                    </TabsContent>

                                    <TabsContent value="accepted" className="m-0">
                                        <div className="divide-y">
                                            {acceptedApplications.length > 0 ? (
                                                acceptedApplications.map((application) => (
                                                    <ApplicationCard
                                                        key={application.id}
                                                        application={application}
                                                        status="accepted"
                                                    />
                                                ))
                                            ) : (
                                                <div className="py-12 text-center">
                                                    <div className="inline-flex h-12 w-12 items-center justify-center rounded-full bg-muted mb-4">
                                                        <Check className="h-6 w-6 text-muted-foreground" />
                                                    </div>
                                                    <h3 className="text-lg font-medium">No accepted applications</h3>
                                                    <p className="text-muted-foreground mt-1 mb-4 max-w-md mx-auto">
                                                        You haven't accepted any applications for this opportunity yet.
                                                    </p>
                                                </div>
                                            )}
                                        </div>
                                    </TabsContent>

                                    <TabsContent value="rejected" className="m-0">
                                        <div className="divide-y">
                                            {rejectedApplications.length > 0 ? (
                                                rejectedApplications.map((application) => (
                                                    <ApplicationCard
                                                        key={application.id}
                                                        application={application}
                                                        status="rejected"
                                                    />
                                                ))
                                            ) : (
                                                <div className="py-12 text-center">
                                                    <div className="inline-flex h-12 w-12 items-center justify-center rounded-full bg-muted mb-4">
                                                        <X className="h-6 w-6 text-muted-foreground" />
                                                    </div>
                                                    <h3 className="text-lg font-medium">No rejected applications</h3>
                                                    <p className="text-muted-foreground mt-1 mb-4 max-w-md mx-auto">
                                                        You haven't rejected any applications for this opportunity yet.
                                                    </p>
                                                </div>
                                            )}
                                        </div>
                                    </TabsContent>
                                </Tabs>
                            </CardContent>
                        </Card>
                    </div>

                    {/* Sidebar with stats */}
                    <div className="space-y-6">
                        <Card>
                            <CardHeader>
                                <CardTitle>Application Stats</CardTitle>
                            </CardHeader>
                            <CardContent className="space-y-6">
                                <div className="space-y-2">
                                    <div className="flex justify-between text-sm">
                                        <span className="text-muted-foreground">Total Applications</span>
                                        <span className="font-medium">{applications.length}</span>
                                    </div>

                                    <div className="space-y-1">
                                        <div className="flex justify-between text-sm">
                                            <span className="flex items-center">
                                                <div className="h-2 w-2 rounded-full bg-blue-500 mr-1"></div>
                                                Pending
                                            </span>
                                            <span>{pendingApplications.length}</span>
                                        </div>
                                        <div className="relative h-2 bg-muted">
                                            <div
                                                className="absolute h-full bg-blue-500"
                                                style={{ width: `${(pendingApplications.length / applications.length) * 100}%` }}
                                            />
                                        </div>
                                    </div>

                                    <div className="space-y-1">
                                        <div className="flex justify-between text-sm">
                                            <span className="flex items-center">
                                                <div className="h-2 w-2 rounded-full bg-green-500 mr-1"></div>
                                                Accepted
                                            </span>
                                            <span>{acceptedApplications.length}</span>
                                        </div>
                                        <div className="relative h-2 bg-muted">
                                            <div
                                                className="absolute h-full bg-green-500"
                                                style={{ width: `${(acceptedApplications.length / applications.length) * 100}%` }}
                                            />
                                        </div>
                                    </div>

                                    <div className="space-y-1">
                                        <div className="flex justify-between text-sm">
                                            <span className="flex items-center">
                                                <div className="h-2 w-2 rounded-full bg-red-500 mr-1"></div>
                                                Rejected
                                            </span>
                                            <span>{rejectedApplications.length}</span>
                                        </div>
                                        <div className="relative h-2 bg-muted">
                                            <div
                                                className="absolute h-full bg-red-500"
                                                style={{ width: `${(rejectedApplications.length / applications.length) * 100}%` }}
                                            />
                                        </div>
                                    </div>
                                </div>
                            </CardContent>
                        </Card>

                        <Card>
                            <CardHeader>
                                <CardTitle>Opportunity Timeline</CardTitle>
                            </CardHeader>
                            <CardContent className="space-y-4">
                                <div className="flex items-center gap-3">
                                    <div className="flex items-center justify-center h-8 w-8 rounded-full bg-muted">
                                        <CalendarClock className="h-4 w-4 text-muted-foreground" />
                                    </div>
                                    <div>
                                        <p className="text-sm font-medium">Created</p>
                                        <p className="text-xs text-muted-foreground">
                                            {format(new Date(opportunity.created_at), "MMMM d, yyyy")}
                                        </p>
                                    </div>
                                </div>

                                <div className="flex items-center gap-3">
                                    <div className="flex items-center justify-center h-8 w-8 rounded-full bg-muted">
                                        <Calendar className="h-4 w-4 text-muted-foreground" />
                                    </div>
                                    <div>
                                        <p className="text-sm font-medium">Starts</p>
                                        <p className="text-xs text-muted-foreground">
                                            {format(new Date(opportunity.start_date), "MMMM d, yyyy")}
                                        </p>
                                    </div>
                                </div>

                                <div className="flex items-center gap-3">
                                    <div className="flex items-center justify-center h-8 w-8 rounded-full bg-muted">
                                        <Clock className="h-4 w-4 text-muted-foreground" />
                                    </div>
                                    <div>
                                        <p className="text-sm font-medium">Ends</p>
                                        <p className="text-xs text-muted-foreground">
                                            {format(new Date(opportunity.end_date), "MMMM d, yyyy")}
                                        </p>
                                    </div>
                                </div>

                                {/* Day countdown */}
                                {new Date(opportunity.start_date) > new Date() && (
                                    <div className="bg-muted/50 p-3 rounded-lg mt-2">
                                        <p className="text-xs text-center text-muted-foreground mb-1">Days until start</p>
                                        <p className="text-xl font-bold text-center">
                                            {Math.ceil((new Date(opportunity.start_date).getTime() - new Date().getTime()) / (1000 * 60 * 60 * 24))}
                                        </p>
                                    </div>
                                )}
                            </CardContent>
                        </Card>

                        <Card>
                            <CardHeader>
                                <CardTitle>Actions</CardTitle>
                            </CardHeader>
                            <CardContent className="space-y-4">
                                <Button className="w-full" onClick={() => navigate(`/edit-project/${projectId}`)}>
                                    Edit Opportunity
                                </Button>
                                <AlertDialog>
                                    <AlertDialogTrigger asChild>
                                        <Button variant="outline" className="w-full">
                                            Cancel Opportunity
                                        </Button>
                                    </AlertDialogTrigger>
                                    <AlertDialogContent>
                                        <AlertDialogHeader>
                                            <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
                                            <AlertDialogDescription>
                                                This will cancel the opportunity and notify all applicants. This action cannot be undone.
                                            </AlertDialogDescription>
                                        </AlertDialogHeader>
                                        <AlertDialogFooter>
                                            <AlertDialogCancel>No, keep opportunity</AlertDialogCancel>
                                            <AlertDialogAction className="bg-destructive text-destructive-foreground hover:bg-destructive/90">
                                                Yes, cancel opportunity
                                            </AlertDialogAction>
                                        </AlertDialogFooter>
                                    </AlertDialogContent>
                                </AlertDialog>
                            </CardContent>
                        </Card>
                    </div>
                </div>
            </div>
        </DashboardLayout>
    );
}

// Application card component
interface ApplicationCardProps {
    application: typeof mockApplications[0];
    onAccept?: (id: number) => void;
    onReject?: (id: number) => void;
    status: "pending" | "accepted" | "rejected";
}

function ApplicationCard({ application, onAccept, onReject, status }: ApplicationCardProps) {
    const [isExpanded, setIsExpanded] = useState(false);

    const getStatusBadge = () => {
        switch (status) {
            case "accepted":
                return <Badge className="bg-green-500 hover:bg-green-600">Accepted</Badge>;
            case "rejected":
                return <Badge variant="destructive">Rejected</Badge>;
            case "pending":
            default:
                return <Badge variant="outline" className="text-blue-500 border-blue-500">Pending</Badge>;
        }
    };

    const getInitials = (name: string) => {
        return name
            .split(' ')
            .map(part => part[0])
            .join('')
            .toUpperCase();
    };

    return (
        <div className="p-4 md:p-6 hover:bg-muted/50 transition-colors">
            <div className="flex flex-col md:flex-row md:items-start gap-4">
                <Avatar className="h-12 w-12">
                    <AvatarImage src={application.volunteer_avatar} alt={application.volunteer_name} />
                    <AvatarFallback>{getInitials(application.volunteer_name)}</AvatarFallback>
                </Avatar>

                <div className="flex-1 space-y-2">
                    <div className="flex flex-col md:flex-row md:items-center justify-between gap-2">
                        <div>
                            <h3 className="font-semibold">{application.volunteer_name}</h3>
                            <p className="text-sm text-muted-foreground">{application.volunteer_email}</p>
                        </div>

                        <div className="flex items-center gap-3">
                            {getStatusBadge()}
                            <div className="text-xs text-muted-foreground">
                                Applied {format(new Date(application.created_at), "MMM d, yyyy")}
                            </div>
                        </div>
                    </div>

                    <div>
                        <div className="flex items-center gap-1 mb-1">
                            <FileText className="h-4 w-4 text-muted-foreground" />
                            <h4 className="text-sm font-medium">Cover Letter</h4>
                        </div>
                        <p className={`text-sm text-muted-foreground ${!isExpanded ? "line-clamp-2" : ""}`}>
                            {application.cover_letter}
                        </p>
                        {application.cover_letter.length > 150 && (
                            <Button
                                variant="link"
                                className="p-0 h-auto text-xs mt-1"
                                onClick={() => setIsExpanded(!isExpanded)}
                            >
                                {isExpanded ? "Show less" : "Show more"}
                            </Button>
                        )}
                    </div>

                    {status === "pending" && onAccept && onReject && (
                        <div className="flex gap-2 mt-4 justify-end">
                            <AlertDialog>
                                <AlertDialogTrigger asChild>
                                    <Button variant="outline" size="sm" className="text-red-500 border-red-200 hover:bg-red-50 hover:text-red-600">
                                        <X className="mr-1 h-4 w-4" />
                                        Reject
                                    </Button>
                                </AlertDialogTrigger>
                                <AlertDialogContent>
                                    <AlertDialogHeader>
                                        <AlertDialogTitle>Reject this application?</AlertDialogTitle>
                                        <AlertDialogDescription>
                                            Are you sure you want to reject {application.volunteer_name}'s application? They will be notified of your decision.
                                        </AlertDialogDescription>
                                    </AlertDialogHeader>
                                    <AlertDialogFooter>
                                        <AlertDialogCancel>Cancel</AlertDialogCancel>
                                        <AlertDialogAction
                                            className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
                                            onClick={() => onReject(application.id)}
                                        >
                                            Reject Application
                                        </AlertDialogAction>
                                    </AlertDialogFooter>
                                </AlertDialogContent>
                            </AlertDialog>

                            <Button size="sm" className="text-green-600 bg-green-50 hover:bg-green-100 hover:text-green-700 border border-green-200" onClick={() => onAccept(application.id)}>
                                <Check className="mr-1 h-4 w-4" />
                                Accept
                            </Button>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
}