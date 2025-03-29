import { useState, useEffect } from "react"
import { useNavigate, useParams } from "react-router-dom"
import { useAuth } from "@/contexts/auth-context"
import { DashboardLayout } from "@/components/dashboard-layout"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"
import { MapPin, CalendarIcon, Clock, ArrowRight, Plus } from "lucide-react"
import { format } from "date-fns"
import { toast, Toaster } from "sonner"
import { z } from "zod"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form"
import { Badge } from "@/components/ui/badge"
import { Skeleton } from "@/components/ui/skeleton"
// Define the opportunity schema for form validation
const opportunityFormSchema = z.object({
    title: z.string().min(5, "Title must be at least 5 characters"),
    description: z.string().min(20, "Description must be at least 20 characters"),
    location: z.string().min(3, "Location is required"),
    date_range: z.object({
        from: z.date({
            required_error: "Start date is required",
        }),
        to: z.date({
            required_error: "End date is required",
        }),
    }).refine(data => data.from <= data.to, {
        message: "End date cannot be before start date",
        path: ["to"],
    }),
})

// Type for the form values
type OpportunityFormValues = z.infer<typeof opportunityFormSchema>

// Mock opportunities for the organization
const mockOrganizationOpportunities = [
    {
        id: 1,
        organization_id: 101,
        title: "Community Garden Clean-up",
        description: "Join us for a weekend clean-up of the community garden. Tasks include weeding, planting new flowers, and general maintenance.",
        location: "Central Park, New York",
        start_date: "2023-06-15",
        end_date: "2023-06-16",
        created_at: "2023-05-20T10:30:00Z",
        updated_at: "2023-05-20T10:30:00Z",
        applications_count: 5
    },
    {
        id: 2,
        organization_id: 101,
        title: "Food Distribution Volunteers",
        description: "Help distribute food packages to families in need. No experience necessary, just a willingness to help!",
        location: "Downtown Community Center",
        start_date: "2023-07-08",
        end_date: "2023-07-09",
        created_at: "2023-06-01T09:45:00Z",
        updated_at: "2023-06-01T09:45:00Z",
        applications_count: 3
    },
    {
        id: 3,
        organization_id: 101,
        title: "Fundraising Event Staff",
        description: "We need volunteers to help with our annual fundraising gala. Roles include greeting guests, managing the silent auction, and general event support.",
        location: "Grand Hotel Ballroom",
        start_date: "2023-08-20",
        end_date: "2023-08-20",
        created_at: "2023-06-15T14:20:00Z",
        updated_at: "2023-06-16T10:15:00Z",
        applications_count: 8
    }
];

type Opportunity = typeof mockOrganizationOpportunities[0];

export function ManageProjectsPage() {
    const { user } = useAuth();
    const navigate = useNavigate();
    const { projectId } = useParams();
    const [opportunities, setOpportunities] = useState<Opportunity[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [selectedTab, setSelectedTab] = useState(projectId ? "view" : "create");

    // Initialize react-hook-form
    const form = useForm<OpportunityFormValues>({
        resolver: zodResolver(opportunityFormSchema),
        defaultValues: {
            title: "",
            description: "",
            location: "",
            date_range: {
                from: undefined,
                to: undefined,
            },
        },
    });

    // Fetch opportunities on component mount
    useEffect(() => {
        const fetchOpportunities = async () => {
            setIsLoading(true);
            try {
                // In a real implementation, this would be an API call
                // const response = await api.get(`/organizations/${user.organizationId}/opportunities`);
                // setOpportunities(response.data);

                // Using mock data for now
                setTimeout(() => {
                    setOpportunities(mockOrganizationOpportunities);
                    setIsLoading(false);
                }, 1000);
            } catch (error) {
                console.error("Error fetching opportunities:", error);
                toast.error("Failed to load opportunities");
                setIsLoading(false);
            }
        };

        fetchOpportunities();
    }, [user]);

    // If a project ID is provided, fetch that specific project
    useEffect(() => {
        if (projectId) {
            setSelectedTab("view");
            // In a real implementation, fetch the specific project details
            // For now, just log that we would fetch it
            console.log(`Would fetch project with ID: ${projectId}`);
        }
    }, [projectId]);

    // Handle form submission
    const onSubmit = async (data: OpportunityFormValues) => {
        try {
            // In a real implementation, this would be an API call
            // const response = await api.post(`/organizations/${user.organizationId}/opportunities`, {
            //   title: data.title,
            //   description: data.description,
            //   location: data.location,
            //   start_date: data.date_range.from,
            //   end_date: data.date_range.to
            // });

            // Simulate API call
            console.log("Form data to submit:", data);

            // Add the new opportunity to the list (normally this would come from the API response)
            const newOpportunity: Opportunity = {
                id: opportunities.length + 1,
                organization_id: 101, // Would be user.organizationId in real implementation
                title: data.title,
                description: data.description,
                location: data.location,
                start_date: data.date_range.from.toISOString(),
                end_date: data.date_range.to.toISOString(),
                created_at: new Date().toISOString(),
                updated_at: new Date().toISOString(),
                applications_count: 0
            };

            setOpportunities([newOpportunity, ...opportunities]);

            // Reset form
            form.reset();

            // Show success message
            toast.success("Opportunity created successfully!");

            // Switch to opportunities tab
            setSelectedTab("view");
        } catch (error) {
            console.error("Error creating opportunity:", error);
            toast.error("Failed to create opportunity");
        }
    };

    // Handle opportunity click
    const handleOpportunityClick = (id: number) => {
        navigate(`/project/${id}`);
    };

    return (
        <DashboardLayout>
            <Toaster position="top-right" richColors />

            <div className="container mx-auto py-6 space-y-6">
                <div className="flex flex-col space-y-2">
                    <h1 className="text-3xl font-bold tracking-tight">Manage Opportunities</h1>
                    <p className="text-muted-foreground">
                        Create and manage volunteer opportunities for your organization
                    </p>
                </div>

                <Tabs value={selectedTab} onValueChange={setSelectedTab} className="space-y-6">
                    <TabsList className="grid w-full grid-cols-2">
                        <TabsTrigger value="create">Create Opportunity</TabsTrigger>
                        <TabsTrigger value="view">View Opportunities</TabsTrigger>
                    </TabsList>

                    <TabsContent value="create">
                        <Card>
                            <CardHeader>
                                <CardTitle>Create New Opportunity</CardTitle>
                                <CardDescription>
                                    Fill out the form below to create a new volunteer opportunity
                                </CardDescription>
                            </CardHeader>
                            <CardContent>
                                <Form {...form}>
                                    <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
                                        <FormField
                                            control={form.control}
                                            name="title"
                                            render={({ field }) => (
                                                <FormItem>
                                                    <FormLabel>Title</FormLabel>
                                                    <FormControl>
                                                        <Input placeholder="Environmental Cleanup" {...field} />
                                                    </FormControl>
                                                    <FormDescription>
                                                        A clear, concise title for your opportunity
                                                    </FormDescription>
                                                    <FormMessage />
                                                </FormItem>
                                            )}
                                        />

                                        <FormField
                                            control={form.control}
                                            name="description"
                                            render={({ field }) => (
                                                <FormItem>
                                                    <FormLabel>Description</FormLabel>
                                                    <FormControl>
                                                        <Textarea
                                                            placeholder="Describe the volunteer work, requirements, and any other relevant details..."
                                                            className="min-h-32"
                                                            {...field}
                                                        />
                                                    </FormControl>
                                                    <FormDescription>
                                                        Provide detailed information about the opportunity
                                                    </FormDescription>
                                                    <FormMessage />
                                                </FormItem>
                                            )}
                                        />

                                        <FormField
                                            control={form.control}
                                            name="location"
                                            render={({ field }) => (
                                                <FormItem>
                                                    <FormLabel>Location</FormLabel>
                                                    <FormControl>
                                                        <Input placeholder="123 Main St, City, State" {...field} />
                                                    </FormControl>
                                                    <FormDescription>
                                                        Where the volunteer work will take place
                                                    </FormDescription>
                                                    <FormMessage />
                                                </FormItem>
                                            )}
                                        />

                                        <FormField
                                            control={form.control}
                                            name="date_range"
                                            render={({ field }) => (
                                                <FormItem className="flex flex-col">
                                                    <FormLabel>Date Range</FormLabel>

                                                    {/* Split into two separate date inputs for better reliability */}
                                                    <div className="grid grid-cols-2 gap-4">
                                                        <div>
                                                            <FormLabel className="text-xs text-muted-foreground">Start Date</FormLabel>
                                                            <Input
                                                                type="date"
                                                                value={field.value?.from ? format(field.value.from, "yyyy-MM-dd") : ""}
                                                                onChange={(e) => {
                                                                    const date = e.target.value ? new Date(e.target.value) : undefined;
                                                                    field.onChange({
                                                                        from: date,
                                                                        to: field.value?.to
                                                                    });
                                                                }}
                                                                className="mt-1"
                                                            />
                                                        </div>

                                                        <div>
                                                            <FormLabel className="text-xs text-muted-foreground">End Date</FormLabel>
                                                            <Input
                                                                type="date"
                                                                value={field.value?.to ? format(field.value.to, "yyyy-MM-dd") : ""}
                                                                onChange={(e) => {
                                                                    const date = e.target.value ? new Date(e.target.value) : undefined;
                                                                    field.onChange({
                                                                        from: field.value?.from,
                                                                        to: date
                                                                    });
                                                                }}
                                                                className="mt-1"
                                                                min={field.value?.from ? format(field.value.from, "yyyy-MM-dd") : ""}
                                                            />
                                                        </div>
                                                    </div>

                                                    {/* Preview of selected dates in a nice format */}
                                                    {field.value?.from && field.value?.to && (
                                                        <div className="mt-2 flex items-center text-sm">
                                                            <CalendarIcon className="mr-2 h-4 w-4 text-muted-foreground" />
                                                            <span>
                                                                {format(field.value.from, "PPP")} - {format(field.value.to, "PPP")}
                                                            </span>
                                                        </div>
                                                    )}

                                                    <FormDescription>
                                                        The start and end dates for this opportunity
                                                    </FormDescription>
                                                    <FormMessage />
                                                </FormItem>
                                            )}
                                        />

                                        <Button type="submit" className="w-full">
                                            Create Opportunity
                                        </Button>
                                    </form>
                                </Form>
                            </CardContent>
                        </Card>
                    </TabsContent>

                    <TabsContent value="view">
                        <div className="grid grid-cols-1 gap-6">
                            {/* Quick create button */}
                            <Button
                                onClick={() => setSelectedTab("create")}
                                className="w-full mb-4 flex items-center justify-center"
                            >
                                <Plus className="mr-2 h-4 w-4" />
                                Create New Opportunity
                            </Button>

                            {isLoading ? (
                                // Loading skeletons
                                Array.from({ length: 3 }).map((_, i) => (
                                    <Card key={i} className="overflow-hidden">
                                        <CardHeader>
                                            <Skeleton className="h-6 w-3/4 mb-2" />
                                            <Skeleton className="h-4 w-1/2" />
                                        </CardHeader>
                                        <CardContent className="space-y-4">
                                            <Skeleton className="h-16 w-full" />
                                            <div className="flex justify-between">
                                                <Skeleton className="h-4 w-1/3" />
                                                <Skeleton className="h-4 w-1/4" />
                                            </div>
                                        </CardContent>
                                        <CardFooter>
                                            <Skeleton className="h-10 w-full" />
                                        </CardFooter>
                                    </Card>
                                ))
                            ) : opportunities.length > 0 ? (
                                // Opportunities list
                                opportunities.map((opportunity) => (
                                    <Card
                                        key={opportunity.id}
                                        className="overflow-hidden cursor-pointer hover:border-primary transition-all"
                                        onClick={() => handleOpportunityClick(opportunity.id)}
                                    >
                                        <CardHeader>
                                            <div className="flex justify-between items-start">
                                                <div>
                                                    <CardTitle>{opportunity.title}</CardTitle>
                                                    <CardDescription className="flex items-center mt-1">
                                                        <MapPin className="h-3.5 w-3.5 mr-1 text-muted-foreground" />
                                                        {opportunity.location}
                                                    </CardDescription>
                                                </div>
                                                <Badge variant="outline" className="ml-2">
                                                    {opportunity.applications_count} {opportunity.applications_count === 1 ? 'application' : 'applications'}
                                                </Badge>
                                            </div>
                                        </CardHeader>
                                        <CardContent className="space-y-4">
                                            <p className="text-sm text-muted-foreground line-clamp-3">
                                                {opportunity.description}
                                            </p>

                                            <div className="flex items-center text-sm text-muted-foreground">
                                                <Clock className="h-4 w-4 mr-1" />
                                                <span>
                                                    {format(new Date(opportunity.start_date), "MMM d, yyyy")} - {format(new Date(opportunity.end_date), "MMM d, yyyy")}
                                                </span>
                                            </div>
                                        </CardContent>
                                        <CardFooter className="flex justify-between">
                                            <span className="text-xs text-muted-foreground">
                                                Created: {format(new Date(opportunity.created_at), "MMM d, yyyy")}
                                            </span>
                                            <Button variant="ghost" size="sm" className="gap-1">
                                                Manage <ArrowRight className="h-4 w-4" />
                                            </Button>
                                        </CardFooter>
                                    </Card>
                                ))
                            ) : (
                                // Empty state
                                <div className="text-center py-12">
                                    <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-muted mb-4">
                                        <Plus className="h-8 w-8 text-muted-foreground" />
                                    </div>
                                    <h3 className="text-lg font-medium">No opportunities yet</h3>
                                    <p className="text-muted-foreground mt-2 mb-6">
                                        Create your first volunteer opportunity to start finding helping hands.
                                    </p>
                                    <Button onClick={() => setSelectedTab("create")}>
                                        Create Your First Opportunity
                                    </Button>
                                </div>
                            )}
                        </div>
                    </TabsContent>
                </Tabs>
            </div>
        </DashboardLayout>
    );
}