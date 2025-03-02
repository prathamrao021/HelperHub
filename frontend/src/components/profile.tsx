// // In your parent component or page
// import { useState, useEffect } from 'react';
// import { Dashboard, DashboardData } from '@components/profile';
// import { useAuth } from '@/contexts/auth-context';

// export function DashboardPage() {
//   const { user } = useAuth();
//   const [dashboardData, setDashboardData] = useState<DashboardData>({
//     user,
//     isLoading: true
//   });

//   useEffect(() => {
//     const fetchDashboardData = async () => {
//       try {
//         // Fetch stat cards
//         const statsResponse = await fetch(`/api/dashboard/stats?userId=${user.id}`);
//         const statsData = await statsResponse.json();
        
//         // Fetch recent jobs
//         const jobsResponse = await fetch(`/api/dashboard/recent-jobs?userId=${user.id}`);
//         const jobsData = await jobsResponse.json();
        
//         // Fetch chart data
//         const chartResponse = await fetch(`/api/dashboard/charts?userId=${user.id}`);
//         const chartData = await chartResponse.json();
        
//         setDashboardData({
//           user,
//           isLoading: false,
//           statCards: statsData.cards,
//           recentJobs: jobsData.jobs,
//           chartData: chartData
//         });
//       } catch (error) {
//         console.error('Error fetching dashboard data:', error);
//         setDashboardData({
//           user,
//           isLoading: false,
//           // Use default values for failed requests
//         });
//       }
//     };
    
//     if (user) {
//       fetchDashboardData();
//     }
//   }, [user]);

//   return <Dashboard dashboardData={dashboardData} />;
// }