import { Dialog, DialogTrigger, DialogContent, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { useState, useEffect } from "react";
import { Form, FormField, FormLabel, FormControl, FormMessage, FormItem } from "@/components/ui/form";
import { User, MapPin, Phone, Clock, Briefcase } from "lucide-react";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useAuth } from "@/contexts/auth-context";
import api from "@/lib/axios";
import { useNavigate } from "react-router-dom";
import { z } from "zod";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { ChevronsUpDown } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { X } from "lucide-react";

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
] as const;

const EditVolunteerFormSchema = z.object({
  name: z
    .string()
    .min(2, "Full name must be at least 2 characters"),
  location: z
    .string()
    .min(3, "Location must be at least 3 characters"),
  phone: z
    .string()
    .min(10, "Phone number must be at least 10 characters"),
  bio_Data: z
    .string()
    .min(50, "Bio must be at least 50 characters")
    .max(500, "Bio must not exceed 500 characters"),
  available_Hours: z
    .number()
    .min(1, "Available hours must be at least 1")
    .max(40, "Available hours must not exceed 40"),
  category_List: z
    .array(z.string())
    .min(1, "Select at least one skill")
});

type VolunteerFormValues = z.infer<typeof EditVolunteerFormSchema>;

export function EditVolunteerProfile() {
  const [open, setOpen] = useState(false);
  const { user, updateUser } = useAuth();
  const navigate = useNavigate();
  
  console.log("Current user data:", user);
  
  const form = useForm<VolunteerFormValues>({
    resolver: zodResolver(EditVolunteerFormSchema),
    defaultValues: {
      name: user?.name || '',
      phone: user?.phone || '',
      location: user?.location || '',
      bio_Data: user?.bio_Data || '',
      available_Hours: user?.available_Hours || 5,
      category_List: user?.category_List || []
    }
  });

  // Update form values when user data changes or dialog opens
  useEffect(() => {
    if (user && open) {
      console.log("Resetting form with user data:", user);
      form.reset({
        name: user.name || '',
        phone: user.phone || '',
        location: user.location || '',
        bio_Data: user.bio_Data || '',
        available_Hours: user.available_Hours || 5,
        category_List: user.category_List || []
      });
    }
  }, [user, form, open]);

  const onSubmit = async (values: VolunteerFormValues) => {
    try {
      const response = await api.put(`volunteers/update/${user?.email}`, {
        Name: values.name,
        Phone: values.phone,
        Location: values.location,
        Bio_Data: values.bio_Data,
        Available_Hours: values.available_Hours,
        Category_List: values.category_List
      });
      
      console.log("Update response:", response.data);
      
      const updatedUser = {
        ...user,
        name: values.name,
        phone: values.phone,
        location: values.location,
        bio_Data: values.bio_Data,
        available_Hours: values.available_Hours,
        category_List: values.category_List
      };

      updateUser(updatedUser);

      setOpen(false);
      window.location.reload();
    } catch (error) {
      console.error("Error updating volunteer profile:", error);
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
          <DialogTitle>Edit Volunteer Profile</DialogTitle>
        </DialogHeader>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Full Name</FormLabel>
                  <FormControl>
                    <div className="relative">
                      <User className="absolute left-3 top-2.5 h-5 w-5 text-muted-foreground" />
                      <Input placeholder="John Doe" className="pl-10" {...field} />
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
                  <FormLabel>Location</FormLabel>
                  <FormControl>
                    <div className="relative">
                      <MapPin className="absolute left-3 top-2.5 h-5 w-5 text-muted-foreground" />
                      <Input 
                        placeholder="City, Country" 
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
              name="bio_Data"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Bio</FormLabel>
                  <FormControl>
                    <Textarea 
                      placeholder="Tell us about yourself, your experience, and what motivates you to volunteer..." 
                      className="min-h-[100px] resize-none"
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
                    <div className="relative">
                      <Clock className="absolute left-3 top-2.5 h-5 w-5 text-muted-foreground" />
                      <Input
                        type="number"
                        min={1}
                        max={40}
                        className="pl-10"
                        {...field}
                        onChange={e => field.onChange(Number(e.target.value))}
                      />
                    </div>
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
                    <div className="space-y-2">
                      <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                          <Button
                            variant="outline"
                            className="w-full justify-between"
                          >
                            <div className="flex items-center gap-1">
                              <Briefcase className="mr-2 h-4 w-4" />
                              {field.value?.length > 0
                                ? `${field.value.length} skill${field.value.length > 1 ? "s" : ""} selected`
                                : "Select skills"}
                            </div>
                            <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                          </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent className="w-full max-h-[200px] overflow-y-auto">
                          <DropdownMenuLabel>Skills</DropdownMenuLabel>
                          <DropdownMenuSeparator />
                          {skills.map((skill) => (
                            <DropdownMenuCheckboxItem
                              key={skill}
                              checked={field.value?.includes(skill)}
                              onCheckedChange={(checked) => {
                                const updatedSkills = checked
                                  ? [...(field.value || []), skill]
                                  : field.value?.filter((s) => s !== skill) || [];
                                field.onChange(updatedSkills);
                              }}
                            >
                              {skill}
                            </DropdownMenuCheckboxItem>
                          ))}
                        </DropdownMenuContent>
                      </DropdownMenu>
                      
                      {/* Display selected skills as badges */}
                      {field.value?.length > 0 && (
                        <div className="flex flex-wrap gap-2 mt-2">
                          {field.value.map((skill) => (
                            <Badge key={skill} variant="secondary" className="px-2 py-1">
                              {skill}
                              <X
                                className="ml-1 h-3 w-3 cursor-pointer"
                                onClick={() => {
                                  field.onChange(
                                    field.value?.filter((s) => s !== skill)
                                  );
                                }}
                              />
                            </Badge>
                          ))}
                        </div>
                      )}
                    </div>
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