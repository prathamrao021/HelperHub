import { Dialog, DialogTrigger, DialogContent, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { useState, useEffect } from "react";
import { Form, FormField, FormLabel, FormControl, FormMessage, FormItem } from "@/components/ui/form";
import { Building, Mail, MapPin, Phone } from "lucide-react";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useAuth } from "@/contexts/auth-context";
import api from "@/lib/axios";
import { useNavigate } from "react-router-dom";
import { z } from "zod";

const EditOrganizationFormSchema = z.object({
  email: z
    .string()
    .email("Please enter a valid email address"),
  name: z
    .string()
    .min(2, "Organization name must be at least 2 characters")
    .max(100, "Organization name must not exceed 100 characters"),
  phone: z.string().min(10, "Phone number must be at least 10 characters"),
  location: z.string().min(10, "Address must be at least 10 characters").max(200, "Address must not exceed 200 characters"),
  description: z.string().min(10, "Description must be at least 10 characters").max(1000, "Description must not exceed 1000 characters"),
});

type OrganizationFormValues = z.infer<typeof EditOrganizationFormSchema>;

export function EditOrganizationProfile() {
  const [open, setOpen] = useState(false);
  const { user, updateUser } = useAuth();
  const navigate = useNavigate();
  
  console.log("Current user data:", user);
  
  const form = useForm<OrganizationFormValues>({
    resolver: zodResolver(EditOrganizationFormSchema),
    defaultValues: {
      email: user?.email || '',
      name: user?.name || '',
      phone: user?.phone || '',
      location: user?.location || '',
      description: user?.description || ''
    }
  });

  // Update form values when user data changes or dialog opens
  useEffect(() => {
    if (user && open) {
      console.log("Resetting form with user data:", user);
      form.reset({
        email: user.email || '',
        name: user.name || '',
        phone: user.phone || '',
        location: user.location || '',
        description: user.description || '',
      });
    }
  }, [user, form, open]);

  const onSubmit = async (values: OrganizationFormValues) => {
    try {
      
      const response = await api.put(`organizations/update/${user?.email}`, {
        Email: values.email,
        Name: values.name,
        Phone: values.phone,
        Location: values.location,
        Description: values.description,
      });
      
      console.log("Update response:", response.data);
      
      const updatedUser = {
        ...user,
        email: values.email,
        name: values.name,
        phone: values.phone,
        location: values.location,
        description: values.description,
      };

      updateUser(updatedUser);

      setOpen(false);
      window.location.reload();
    } catch (error) {
      console.error("Error updating organization profile:", error);
    }
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <div>
          <Button onClick={() => setOpen(true)}>Edit Profile</Button>
        </div>
      </DialogTrigger>

      <DialogContent className="max-w-lg max-h-[80vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Edit Organization Profile</DialogTitle>
        </DialogHeader>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <div className="relative">
                      <Mail className="absolute left-3 top-2.5 h-5 w-5 text-muted-foreground" />
                      <Input placeholder="you@example.com" className="pl-10" disabled {...field} />
                    </div>
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Organization Name</FormLabel>
                  <FormControl>
                    <div className="relative">
                      <Building className="absolute left-3 top-2.5 h-5 w-5 text-muted-foreground" />
                      <Input placeholder="Your Organization Name" className="pl-10" {...field} />
                    </div>
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
                    <div className="relative">
                      <Phone className="absolute left-3 top-2.5 h-5 w-5 text-muted-foreground" />
                      <Input placeholder="+1 (555) 000-0000" className="pl-10" {...field} />
                    </div>
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
                  <FormLabel>Address</FormLabel>
                  <FormControl>
                    <div className="relative">
                      <MapPin className="absolute left-3 top-2.5 h-5 w-5 text-muted-foreground" />
                      <Input 
                        placeholder="123 Main St, City, Country" 
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
              name="description"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Organization Description</FormLabel>
                  <FormControl>
                    <Textarea 
                      placeholder="Tell us about your organization's mission, goals, and the impact you want to make..." 
                      className="min-h-[150px] resize-none"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <Button type="submit" className="w-full">
              Save Changes
            </Button>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}