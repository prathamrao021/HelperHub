import { useState } from "react";
import { Container, Tabs, Tab, Box, Card, CardContent, Typography } from "@mui/material";


type ProfileType = "volunteer" | "organization";


const volunteerData = {
  name: "John Doe",
  skills: ["Gardening", "Grocery Assistance", "Pet Care"],
  tasksCompleted: ["Planted a community garden - Jan 2024", "Delivered groceries to seniors - Dec 2023"],
};

const organizationData = {
  name: "Helping Hands",
  mission: "To empower communities through volunteer work and social initiatives.",
  ongoingOpportunities: ["Community Cleanup", "Food Distribution Drive"],
  contact: "contact@helpinghands.org",
};

export default function Profile() {
  const [profileType, setProfileType] = useState<ProfileType>("volunteer");

  return (
    <Container maxWidth="md">
      <Tabs value={profileType} onChange={(_, newValue) => setProfileType(newValue)}>
        <Tab label="Volunteer Profile" value="volunteer" />
        <Tab label="Organization Profile" value="organization" />
      </Tabs>

      <Box mt={3}>
        {profileType === "volunteer" ? <VolunteerProfile /> : <OrganizationProfile />}
      </Box>
    </Container>
  );
}

function VolunteerProfile() {
  return (
    <Box>
      <ProfileCard title="Personal Info">
        <Typography>Name: {volunteerData.name}</Typography>
      </ProfileCard>

      <ProfileCard title="Skills Showcase">
        {volunteerData.skills.map((skill, index) => (
          <Typography key={index}>- {skill}</Typography>
        ))}
      </ProfileCard>

      <ProfileCard title="Tasks Completed">
        {volunteerData.tasksCompleted.map((task, index) => (
          <Typography key={index}>- {task}</Typography>
        ))}
      </ProfileCard>
    </Box>
  );
}

function OrganizationProfile() {
  return (
    <Box>
      <ProfileCard title="Mission Statement">
        <Typography>{organizationData.mission}</Typography>
      </ProfileCard>

      <ProfileCard title="Ongoing Opportunities">
        {organizationData.ongoingOpportunities.map((opportunity, index) => (
          <Typography key={index}>- {opportunity}</Typography>
        ))}
      </ProfileCard>

      <ProfileCard title="Contact Info">
        <Typography>Email: {organizationData.contact}</Typography>
      </ProfileCard>
    </Box>
  );
}


function ProfileCard({ title, children }: { title: string; children: React.ReactNode }) {
  return (
    <Card sx={{ mb: 2, p: 2 }}>
      <CardContent>
        <Typography variant="h6">{title}</Typography>
        {children}
      </CardContent>
    </Card>
  );
}
