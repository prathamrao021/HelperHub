import React from "react";
import { Bar, BarChart, ResponsiveContainer, XAxis, YAxis, Tooltip, Legend } from "recharts";

// Mock data for the charts
const volunteerData = [
  { name: "Jan", hours: 8 },
  { name: "Feb", hours: 16 },
  { name: "Mar", hours: 12 },
  { name: "Apr", hours: 22 },
  { name: "May", hours: 30 },
  { name: "Jun", hours: 20 },
  { name: "Jul", hours: 24 },
];

const organizationData = [
  { name: "Jan", volunteers: 10, hours: 45 },
  { name: "Feb", volunteers: 15, hours: 72 },
  { name: "Mar", volunteers: 12, hours: 56 },
  { name: "Apr", volunteers: 18, hours: 90 },
  { name: "May", volunteers: 22, hours: 110 },
  { name: "Jun", volunteers: 19, hours: 95 },
  { name: "Jul", volunteers: 25, hours: 125 },
];

interface OverviewProps {
  userRole: string;
}

const Overview: React.FC<OverviewProps> = ({ userRole }) => {
  const isOrganization = userRole === "ORGANIZATION_ADMIN";
  
  if (isOrganization) {
    return (
      <ResponsiveContainer width="100%" height={300}>
        <BarChart data={organizationData}>
          <XAxis 
            dataKey="name" 
            stroke="#888888" 
            fontSize={12} 
            tickLine={false} 
            axisLine={false}
          />
          <YAxis 
            stroke="#888888" 
            fontSize={12} 
            tickLine={false} 
            axisLine={false}
            tickFormatter={(value) => `${value}`}
          />
          <Tooltip 
            content={({ active, payload }) => {
              if (active && payload && payload.length) {
                return (
                  <div className="rounded-lg border bg-background p-2 shadow-sm">
                    <div className="grid grid-cols-2 gap-2">
                      <div className="flex flex-col">
                        <span className="text-[0.70rem] uppercase text-muted-foreground">
                          Volunteers
                        </span>
                        <span className="font-bold text-foreground">
                          {payload[0].value}
                        </span>
                      </div>
                      <div className="flex flex-col">
                        <span className="text-[0.70rem] uppercase text-muted-foreground">
                          Hours
                        </span>
                        <span className="font-bold text-foreground">
                          {payload[1].value}
                        </span>
                      </div>
                    </div>
                  </div>
                );
              }
              return null;
            }}
          />
          <Legend />
          <Bar 
            dataKey="volunteers" 
            fill="var(--chart-1)" 
            radius={4} 
            className="fill-primary" 
          />
          <Bar 
            dataKey="hours" 
            fill="var(--chart-2)" 
            radius={4} 
            className="fill-secondary" 
          />
        </BarChart>
      </ResponsiveContainer>
    );
  }
  
  return (
    <ResponsiveContainer width="100%" height={300}>
      <BarChart data={volunteerData}>
        <XAxis 
          dataKey="name" 
          stroke="#888888" 
          fontSize={12} 
          tickLine={false} 
          axisLine={false}
        />
        <YAxis 
          stroke="#888888" 
          fontSize={12} 
          tickLine={false} 
          axisLine={false}
          tickFormatter={(value) => `${value}`}
        />
        <Tooltip 
          content={({ active, payload }) => {
            if (active && payload && payload.length) {
              return (
                <div className="rounded-lg border bg-background p-2 shadow-sm">
                  <div className="flex flex-col">
                    <span className="text-[0.70rem] uppercase text-muted-foreground">
                      Hours
                    </span>
                    <span className="font-bold text-foreground">
                      {payload[0].value}
                    </span>
                  </div>
                </div>
              );
            }
            return null;
          }}
        />
        <Bar 
          dataKey="hours" 
          fill="var(--chart-1)" 
          radius={4} 
          className="fill-primary" 
        />
      </BarChart>
    </ResponsiveContainer>
  );
};

export default Overview;