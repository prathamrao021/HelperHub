import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { useNavigate } from "react-router-dom"
import { Navbar } from "@/components/Navbar"
import { z } from "zod"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { ChevronsUpDown } from "lucide-react"
import { Badge } from "@/components/ui/badge"
import { X } from "lucide-react"
import { useAuth } from "@/contexts/auth-context"

const skills = [
  "Web Development",
  "Graphic Design",
  "Content Writing",
  "Social Media",
  "Teaching",
  "Event Planning",
  "Photography",
  "Translation",
  "First Aid",
  "Project Management",
  "Other"
] as const

const volunteerFormSchema = z.object({
  name: z.string().min(2, "Full name must be at least 2 characters"),
  email: z.string().email("Please enter a valid email address"),
  password: z.string().min(6, "Password must be at least 6 characters"),
  phone: z.string().min(10, "Phone number must be at least 10 digits"),
  bio_Data: z.string().min(50, "Bio must be at least 50 characters").max(500, "Bio must not exceed 500 characters"),
  available_Hours: z.number().min(1).max(40),
  location: z.string().min(3, "Location must be at least 3 characters"),
  category_List: z.array(z.string()).min(1, "Select at least one skill")
})

type VolunteerFormValues = z.infer<typeof volunteerFormSchema>

export function VolunteerRegistration() {
  const navigate = useNavigate()
  const { registerVolunteer } = useAuth()

  const form = useForm<VolunteerFormValues>({
    resolver: zodResolver(volunteerFormSchema),
    defaultValues: {
      name: "",
      email: "",
      password: "",
      phone: "",
      bio_Data: "", 
      location: "",
      category_List: [],
      available_Hours: 5
    }
  })

  function onSubmit(data: VolunteerFormValues) {
    registerVolunteer(data)
      .then(() => navigate("/dashboard"))
      .catch((error) => console.error("Error registering volunteer:", error))
  }

  return (
    <>
      <Navbar
        menuItems={[]}
        showThemeToggle={true}
        showLoginButton={true}
        onLoginClick={() => navigate("/login")}
      />
      <div className="flex min-h-screen w-full items-center justify-center p-6 md:p-10">
        <div className="container max-w-2xl mx-auto px-4 py-8">
          <Card className="backdrop-blur-sm border-muted/40">
            <CardHeader className="space-y-1">
              <CardTitle className="text-2xl font-bold">Complete Your Volunteer Profile</CardTitle>
              <CardDescription>
                Tell us more about yourself and your volunteering interests
              </CardDescription>
            </CardHeader>
            <CardContent>
              <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
                  <FormField
                    control={form.control}
                    name="name"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Full Name</FormLabel>
                        <FormControl>
                          <Input placeholder="John Doe" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="email"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Email Address</FormLabel>
                        <FormControl>
                          <Input placeholder="abcd@gmail.com" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="password"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Password</FormLabel>
                        <FormControl>
                          <Input type="password" placeholder="••••••••" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="phone"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Phone Number</FormLabel>
                        <FormControl>
                          <Input placeholder="+1 (555) 000-0000" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <FormField
                    control={form.control}
                    name="bio_Data"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Bio</FormLabel>
                        <FormControl>
                          <Textarea
                            placeholder="Tell us about yourself, your experience, and what motivates you to volunteer..."
                            className="min-h-[100px]"
                            {...field}
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <FormField
                    control={form.control}
                    name="available_Hours"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Weekly Available Hours</FormLabel>
                        <FormControl>
                          <Input
                            type="number"
                            min={1}
                            max={40}
                            {...field}
                            onChange={e => field.onChange(Number(e.target.value))}
                          />
                        </FormControl>
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
                          <Input placeholder="City, Country" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <FormField
                    control={form.control}
                    name="category_List"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Skills</FormLabel>
                        <FormControl>
                          <DropdownMenu>
                            <DropdownMenuTrigger>
                              <Button
                                variant="outline"
                                className="w-full justify-between"
                              >
                                {field.value?.length > 0
                                  ? `${field.value.length} skill${field.value.length > 1 ? "s" : ""} selected`
                                  : "Select skills"}
                                <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                              </Button>
                            </DropdownMenuTrigger>
                            <DropdownMenuContent className="w-56">
                              <DropdownMenuLabel>Available Skills</DropdownMenuLabel>
                              <DropdownMenuSeparator />
                              {skills.map((skill) => (
                                <DropdownMenuCheckboxItem
                                  key={skill}
                                  checked={field.value?.includes(skill)}
                                  onCheckedChange={(checked) => {
                                    const currentValue = field.value || []
                                    const newValue = checked
                                      ? [...currentValue, skill]
                                      : currentValue.filter((s) => s !== skill)
                                    field.onChange(newValue)
                                  }}
                                >
                                  {skill}
                                </DropdownMenuCheckboxItem>
                              ))}
                            </DropdownMenuContent>
                          </DropdownMenu>

                        </FormControl>
                        <div className="flex flex-wrap gap-2 mt-2">
                          {field.value?.map((skill) => (
                            <Badge
                              key={skill}
                              variant="secondary"
                              className="flex items-center gap-1 px-2 py-1"
                            >
                              {skill}
                              <X
                                className="h-3 w-3 cursor-pointer"
                                onClick={() => {
                                  const newValue = field.value?.filter((s) => s !== skill);
                                  field.onChange(newValue);
                                }}
                              />
                            </Badge>
                          ))}
                        </div>
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <Button type="submit" className="w-full">
                    Complete Profile
                  </Button>
                </form>
              </Form>
            </CardContent>
          </Card>
        </div>
      </div>
    </>
  )
}