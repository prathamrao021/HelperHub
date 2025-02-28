import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { UserCircle2, Mail, Lock, Users } from "lucide-react"
import { Navbar } from "@/components/Navbar"

const userRegistrationSchema = z.object({
    fullName: z
        .string()
        .min(2, { message: "Full name must be at least 2 characters." })
        .max(100, { message: "Full name must be at most 100 characters." }),
    email: z.string().email({ message: "Please enter a valid email address." }),
    password: z
        .string()
        .min(6, { message: "Password must be at least 6 characters." }),
    userRole: z.enum(["VOLUNTEER", "ORGANIZATION_ADMIN"]).default("VOLUNTEER"),
})

type UserRegistrationFormValues = z.infer<typeof userRegistrationSchema>

export function Register() {
    const form = useForm<UserRegistrationFormValues>({
        resolver: zodResolver(userRegistrationSchema),
        defaultValues: {
            userRole: "VOLUNTEER",
        },
        mode: "onChange",
    })

    function onSubmit(data: UserRegistrationFormValues) {
        console.log(data)
    }

    return (
        <>
            <Navbar
                menuItems={[]}
                showThemeToggle={true}
                showLoginButton={true}
                onLoginClick={() => console.log("Login button clicked")}
            />
            <div className="flex min-h-screen w-full items-center justify-center p-6 md:p-10">
                <div className="container max-w-md mx-auto px-4 py-8">
                    <Card className="backdrop-blur-sm border-muted/40">
                        <CardHeader className="space-y-1">
                            <CardTitle className="text-2xl font-bold">Create an account</CardTitle>
                            <CardDescription>
                                Enter your details below to create your account
                            </CardDescription>
                        </CardHeader>
                        <CardContent>
                            <Form {...form}>
                                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
                                    <FormField
                                        control={form.control}
                                        name="fullName"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>Full Name</FormLabel>
                                                <FormControl>
                                                    <div className="relative">
                                                        <UserCircle2 className="absolute left-3 top-2.5 h-5 w-5 text-muted-foreground" />
                                                        <Input placeholder="John Doe" className="pl-10" {...field} />
                                                    </div>
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
                                                <FormLabel>Email</FormLabel>
                                                <FormControl>
                                                    <div className="relative">
                                                        <Mail className="absolute left-3 top-2.5 h-5 w-5 text-muted-foreground" />
                                                        <Input placeholder="you@example.com" className="pl-10" {...field} />
                                                    </div>
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
                                                    <div className="relative">
                                                        <Lock className="absolute left-3 top-2.5 h-5 w-5 text-muted-foreground" />
                                                        <Input
                                                            type="password"
                                                            placeholder="••••••••"
                                                            className="pl-10"
                                                            {...field}
                                                        />
                                                    </div>
                                                </FormControl>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />

                                    <FormField
                                        control={form.control}
                                        name="userRole"
                                        render={({ field }) => (
                                            <FormItem>
                                                <FormLabel>Role</FormLabel>
                                                <Select onValueChange={field.onChange} defaultValue={field.value}>
                                                    <FormControl>
                                                        <div className="relative">
                                                            <Users className="absolute left-3 top-2.5 h-5 w-5 text-muted-foreground" />
                                                            <SelectTrigger className="pl-10">
                                                                <SelectValue placeholder="Select a role" />
                                                            </SelectTrigger>
                                                        </div>
                                                    </FormControl>
                                                    <SelectContent>
                                                        <SelectItem value="VOLUNTEER">Volunteer</SelectItem>
                                                        <SelectItem value="ORGANIZATION_ADMIN">Organization Admin</SelectItem>
                                                    </SelectContent>
                                                </Select>
                                                <FormMessage />
                                            </FormItem>
                                        )}
                                    />

                                    <Button type="submit" className="w-full">Create Account</Button>

                                    <div className="text-center text-sm text-muted-foreground">
                                        Already have an account?{" "}
                                        <a href="/login" className="text-primary underline-offset-4 hover:underline">
                                            Sign in
                                        </a>
                                    </div>
                                </form>
                            </Form>
                        </CardContent>
                    </Card>
                </div>
            </div>

        </>
    )
}