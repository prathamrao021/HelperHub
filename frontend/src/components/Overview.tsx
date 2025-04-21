import React from "react"
import { Bar, BarChart, ResponsiveContainer, XAxis, YAxis, Tooltip, Legend } from "recharts"
import { ChartContainer } from "@/components/ui/chart"
import type { ChartConfig } from "@/components/ui/chart" // Import the type
import { Circle } from "lucide-react";

// Mock data for the charts
const volunteerData = [
  { name: "Jan", hours: 12 },
  { name: "Feb", hours: 19 },
  { name: "Mar", hours: 15 },
  { name: "Apr", hours: 22 },
  { name: "May", hours: 28 },
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

// Create chart configurations
// Create correct chart configurations
const organizationConfig: ChartConfig = {
    // Based on the error, it seems the config only accepts legend properties
    // Remove the y property that's causing issues
    legend: {
      label: "Organization Data",
      icon: Circle
    }
  };
  
  const volunteerConfig: ChartConfig = {
    // Only include valid properties
    legend: {
      label: "Volunteer Hours",
      icon: Circle
    }
  };

  const Overview: React.FC<{ userRole: string }> = ({ userRole }) => {
    const isOrganization = userRole === "ORGANIZATION_ADMIN";
    
    if (isOrganization) {
      return (
        <div className="h-[240px]">
          <ChartContainer config={organizationConfig}>
            <ResponsiveContainer width="100%" height="100%">
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
                  domain={[0, 130]} // Set min/max directly on YAxis instead of in config
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
                              <span className="font-bold text-muted-foreground">
                                {payload[0].value}
                              </span>
                            </div>
                            <div className="flex flex-col">
                              <span className="text-[0.70rem] uppercase text-muted-foreground">
                                Hours
                              </span>
                              <span className="font-bold text-muted-foreground">
                                {payload[1].value}
                              </span>
                            </div>
                          </div>
                        </div>
                      )
                    }
                    return null
                  }}
                />
                <Legend />
                <Bar 
                  dataKey="volunteers" 
                  fill="hsl(var(--primary))" 
                  radius={4} 
                  className="fill-primary" 
                />
                <Bar 
                  dataKey="hours" 
                  fill="hsl(var(--secondary))" 
                  radius={4} 
                  className="fill-secondary" 
                />
              </BarChart>
            </ResponsiveContainer>
          </ChartContainer>
        </div>
      )
    }
    
    return (
      <div className="h-[240px]">
        <ChartContainer config={volunteerConfig}>
          <ResponsiveContainer width="100%" height="100%">
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
                domain={[0, 30]} // Set min/max directly on YAxis instead of in config
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
                          <span className="font-bold text-muted-foreground">
                            {payload[0].value}
                          </span>
                        </div>
                      </div>
                    )
                  }
                  return null
                }}
              />
              <Bar 
                dataKey="hours" 
                fill="hsl(var(--primary))" 
                radius={4} 
                className="fill-primary" 
              />
            </BarChart>
          </ResponsiveContainer>
        </ChartContainer>
      </div>
    )
  }
  
export default Overview;