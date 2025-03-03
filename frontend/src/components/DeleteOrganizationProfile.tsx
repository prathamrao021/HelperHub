import { Dialog, DialogTrigger, DialogContent, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { useState, useEffect } from "react";
import { Building, Mail, MapPin, Phone } from "lucide-react";
import { useAuth } from "@/contexts/auth-context";
import api from "@/lib/axios";
import { useNavigate } from "react-router-dom";


export function DeleteOrganizationProfile() {
  const [open, setOpen] = useState(false);
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <div>
          <Button className="ml-5" variant="destructive"onClick={() => setOpen(true)}>Delete Profile</Button>
        </div>
      </DialogTrigger>

      <DialogContent className="max-w-lg max-h-[80vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Are you Sure????</DialogTitle>
            <div className="mt-4">
                <Button variant="destructive" className="mr-4" onClick={() => {
                    api.delete(`organizations/delete/${user?.email}`)
                    .then(() => {
                        navigate("/");
                        logout();
                    })
                    .catch((error) => {
                        console.error("Error deleting organization profile:", error);
                    });
                }}>Delete</Button>
                <Button onClick={() => setOpen(false)}>Cancel</Button>
            </div>
        </DialogHeader>
        
      </DialogContent>
    </Dialog>
  );
}