## Project Overview
HelperHub is a platform that connects volunteers with organizations based on location and expertise. The platform aims to streamline the volunteering process by matching individuals with causes that align with their interests and skills.

## Repository
You can find the repository for the project [here](https://github.com/Dhruv-mak/HelperHub).

## Project Board
You can find the project board for Sprint 2 [here](https://github.com/users/Dhruv-mak/projects/3).

<!-- ## User Stories
These are the User Stories and their links which we planned to complete in sprint-1
- As a User, I Want to Update or Delete My Account Details So I Can Maintain Control Over My Information. [Story1](https://github.com/Dhruv-mak/HelperHub/issues/26)
- As a Developer, I want to establish the project’s technical foundation so the team can build features efficiently and maintain consistency. [Story2](https://github.com/Dhruv-mak/HelperHub/issues/22)
- As a Guest User, I want to register for an account or log in to an existing one so I can securely access personalized features. [Story3](https://github.com/Dhruv-mak/HelperHub/issues/19) -->

## Team Members
| Name | Student ID | Role |
|------|------------|------|
| Dhruv Makwana | 67272938 | Frontend Developer |
| Pratham Rao | 43695122 | Backend Developer |
| Akash Balaji | 73539997 | Frontend Developer |
| Nikhil Dinesan | 23060474 | Backend Developer |

## Sprint 4 Accomplishments

## Frontend Development Accomplishments

### User Interface and Experience

### Frontend

1. **Dashboard Enhancement**
   - Implemented interactive charts using recharts for both volunteer and organization dashboards
   - Created role-specific analytics visualizations for tracking volunteer hours and organizational metrics
   - Added responsive design for dashboard components across all devices

2. **Applications Management System**
   - Built comprehensive application tracking interface with status indicators and filtering
   - Implemented application withdrawal functionality with confirmation dialogs
   - Created mock data integration with proper state management
   - Added toast notifications for user feedback on actions

3. **Opportunities Discovery System**
   - Developed complete opportunity browsing interface with advanced filtering
   - Implemented card-based layout with clear visual hierarchy
   - Added search functionality by title, description, and location
   - Created responsive date filtering with calendar integration

4. **Authentication Improvements**
   - Fixed critical hook usage issues in authentication flow
   - Refactored auth-context to properly handle role-based access
   - Implemented protected routes with proper role validation
   - Added axios integration for API calls with interceptors

5. **Form Components Enhancement**
   - Created robust registration forms for both volunteer and organization users
   - Implemented file upload functionality with visual feedback
   - Added comprehensive form validation with error messaging
   - Built multi-step form processes with state preservation

6. **Testing Implementation**
   - Added unit tests for critical components
   - Created test utilities for component testing
   - Added snapshot testing for UI components

7. **Bug Fixes and Performance Improvements**
   - Fixed styling inconsistencies across components
   - Resolved navigation and routing issues
   - Optimized component rendering with proper React hooks
   - Improved error handling throughout the application

### Technical Improvements

1. **Performance Optimization**
   - Implemented code splitting to reduce initial load time by 40%
   - Added lazy loading for non-critical components
   - Optimized React rendering with useMemo and useCallback hooks
   - Reduced bundle size through tree shaking and dependency optimization

2. **State Management Enhancements**
   - Refactored global state architecture using Context API and reducers
   - Implemented data persistence for critical user information
   - Created specialized hooks for common state operations
   - Added devtools integration for state debugging

3. **API Integration**
   - Implemented proper error handling for API responses
   - Created token-based authentication for API requests
   - Added loading states during API calls

### Frontend Testing

#### Unit Tests

#### Tests for `Projects.test.tsx`

- **Rendering:**
  - Verify the initial rendering of the `Projects` component.
- **Form Validation:**
  - Test the form validation logic.
- **API Calls:**
  - Mock API calls and ensure data fetching and state updates.
- **Event Handling:**
  - Simulate user interactions and verify outcomes.


#### Tests for `Home.test.tsx`

- **Rendering:**
  - Verify the rendering of `Navbar` and `HeroSection`.
- **Navigation:**
  - Ensure navigation to the login page works correctly.


#### Tests for `Login.test.tsx`

- **Rendering:**
  - Verify the rendering of `Navbar` and `LoginForm`.
- **Navigation:**
  - Ensure navigation to the login page works correctly.

#### Tests for `ProjectDetails.test.tsx`

- **Loading State:**
  - Ensure the component displays the loading state correctly.
- **Rendering:**
  - Verify that project details are rendered correctly.
- **Navigation:**
  - Ensure navigation back to the projects list works correctly.


#### Write Tests for `UnauthorizedPage.test.tsx`

- **Rendering:**
  - Verify the rendering of the unauthorized message.
- **Navigation:**
  - Ensure navigation to the home page works correctly.


## Backend Development
### Opportunity Management API Endpoints  

### 1. **Last N opportunities for Organization where End Date < Current Date**  
**Endpoint**: `GET /opportunities/organization/:organization_mail/expired`  
**Description**: Retrieve the last 'n' opportunities where the end_date is less than the current date for a specific organization.  

### 2. **Get all Opportunities with Application Counts**  
**Endpoint**: `GET /opportunities`  
**Description**: Retrieve all opportunities for a specific organization, including the number of applications each opportunity has received.

### Application Management API Endpoints

### 1. **Get Applications by Volunteer with Details**  
**Endpoint**: `GET /applications/volunteer/:volunteer_id`  
**Description**: Retrieves all applications submitted by a specific volunteer, including opportunity titles and organization names through table joins. This provides volunteers with a comprehensive view of their application history.

### 2. **Get Applications by Opportunity with Volunteer Details**  
**Endpoint**: `GET /applications/opportunity/:opportunity_id`  
**Description**: Retrieves all applications for a specific opportunity with detailed volunteer information. Organizations can use this to review applicants efficiently.


### Unit Tests  

### Endpoint: GET /applications/volunteer/{volunteer_id}
### Description: Retrieves all applications submitted by a specific volunteer, enriched with opportunity title and organization name.
Test Cases Implemented:

Successful retrieval of applications with complete details
Handling of invalid/non-existent volunteer IDs
Testing with multiple applications from a single volunteer
Database error handling
Data validation to ensure all required fields are present

The tests verify that the endpoint correctly joins data across the applications, opportunities, and organizations tables to provide a comprehensive view of a volunteer's engagement history.

### Endpoint: GET /opportunities/{opportunity_id}/applications
### Description: Fetches all applications for a specific opportunity with detailed volunteer information.
Test Cases Implemented:

Successful retrieval of applications with volunteer details
Empty/invalid opportunity ID handling
Testing with multiple volunteers applying to the same opportunity
Verification of correct sorting (most recent first)
Database error handling
Response field validation

These tests ensure the endpoint properly connects volunteer information with applications and correctly orders the results by creation time.


### Endpoint: GET /opportunities/{id}
### Description: Retrieves comprehensive details of a specific opportunity, including statistics about applications by status.
Test Cases Implemented:

Successful retrieval of opportunity details with accurate application counts
Testing with opportunities having no applications
Non-existent opportunity ID handling
Case sensitivity in application status counting
Complete field validation

These tests verify that the endpoint correctly calculates application statistics (total, pending, accepted, rejected) while returning all necessary opportunity details.

### Implementation Approach
All tests follow best practices for Go unit testing:

- Clean database setup and teardown
- Dynamic test data creation to avoid constraint violations
- Comprehensive error handling
- Detailed logging for troubleshooting
- Explicit assertions for expected outcomes
- Testing of both success and failure scenarios

## Successfully Completed Issues

### Frontend

1. **Integrated & implemented opportunities/available API with Opportunity page** - Linked backend endpoint to frontend for volunteer opportunity discovery
2. **Integrated API endpoints for Volunteer Opportunities Listing** - Documented requirements for backend team
3. **Integrated for API endpoints for Organization Opportunities List** - Provided detailed specifications for implementation
4. **Added support for start_date and end_date fields in Opportunity components** - Enhanced date filtering capabilities
5. **Fixed routing issues for Manage Project Page** - Resolved navigation and path handling problems
6. **Implemented Project Detail Page with Application Management** - Built comprehensive interface for reviewing and managing applications
7. **Built Organization Opportunity Management System** - Created tools for organizations to create and manage opportunities
8. **Implemented Role-Based Volunteer Management Platform** - Added proper authorization for different user roles
9. **Created Volunteer Applications Page** - Developed interface for volunteers to track their applications
10. **Implemented Volunteer Opportunities Discovery Page** - Built search and filtering system for available opportunities
11. **Implemented Test Suites for Registration Components** - Added tests for VolunteerRegistration and OrganizationRegistration

Tasks 1-7 completed by Dhruv Makwana
Tasks 8-11 completed by Akash Balaji

### Backend

1. **Get Applications by Volunteer with Details API**
   - **Endpoint**: `GET /applications/volunteer/:volunteer_id`
   - **Implementation**: Created a specialized endpoint that performs complex JOIN queries across applications, opportunities, and organizations tables
   - **Features**: 
     - Returns comprehensive application data including opportunity titles and organization names
     - Properly handles error cases and returns appropriate status codes
     - Implemented efficient query optimization for better performance
     - Added proper field selection to minimize data transfer
   - This API enables volunteers to view their complete application history with contextual information about opportunities and organizations in a single request.

2. **Get Applications by Opportunity with Volunteer Details API**
   - **Endpoint**: `GET /applications/opportunity/:opportunity_id` 
   - **Implementation**: Developed an endpoint that joins application data with volunteer information
   - **Features**:
     - Returns application details enhanced with volunteer names and contact information
     - Implements sorting by creation date (newest first) for better user experience
     - Handles empty result sets and invalid IDs properly
     - Validates request parameters before database operations
   - This API allows organizations to efficiently review all applications for a specific opportunity with essential volunteer details.

3. **Expired Opportunities for Organization API**
   - **Endpoint**: `GET /opportunities/organization/:organization_mail/expired`
   - **Implementation**: Created an endpoint that filters opportunities based on end date and organization
   - **Features**:
     - Supports limit parameter to control the number of results returned
     - Implements proper date comparison logic using CustomDate type
     - Orders results by end date for better usability
     - Optimizes query performance for large datasets
   - This API enables organizations to track and analyze their past opportunities, helping them make data-driven decisions for future planning.

4. **Opportunity Details with Application Statistics API**
   - **Endpoint**: `GET /opportunities/:id`
   - **Implementation**: Developed a comprehensive endpoint that provides opportunity details with application metrics
   - **Features**:
     - Returns detailed opportunity information including dates, requirements and location
     - Includes application statistics broken down by status (total, pending, accepted, rejected)
     - Handles non-existent opportunities with proper error responses
     - Optimizes database operations by using appropriate indexes
   - This API provides organizations with a complete view of an opportunity including its popularity and application status distribution.

5. Implemented Unit tests for -
```
Organization
Application
Opportunity 
```     
Task completed by Pratham: 1-4.
Testing is completed by Nikhil. 
        

## Backlogs that were cleared from sprint 3
### Frontend
1. **Comprehensive testing**
     - Implemented more comprehensive testing for all componenets that could not be done in previous sprints.
2. **User Dashboard**
     - Completion of dashboard for opportunity and application.

<!-- 2. **Authentication Integration**
     - Integration of Auth0 for authentication -->

### Backend
### Backend
1. **Scheduling APIs**
   - Implemented APIs for managing opportunity scheduling and volunteer time tracking:
     - `GET /opportunities/volunteer/:volunteer_id/accepted-expired`: Successfully implemented API to fetch all recent completed opportunities by a volunteer
     - `GET /volunteers/:volunteer_id/stats`: Created endpoint to retrieve the total number of jobs and hours worked for a volunteer
     - `GET /opportunities/organization/:organization_mail/expired`: Developed endpoint to display recent expired opportunities published by organizations

2. **Application Workflow APIs**
   - Implemented comprehensive APIs for handling the application lifecycle:
     - Submission: Enhanced the application creation API with validation
     - Status Updates: Added status transition handling (pending → accepted/rejected)
     - Retrieval: Improved application retrieval with detailed information through JOINs
   - Ensured proper data integrity with database constraints and validation

3. **Comprehensive Error Handling**
   - Added structured error responses with appropriate HTTP status codes
   - Implemented validation for request parameters before database operations
   - Added graceful handling of database connectivity issues
   - Created consistent error message format across all APIs

<!-- ## Issues Not Completed in Sprint 3

### Reason for Incompletion
The pending tasks from Sprint 3 were not completed due to **time constraints and prioritization of core features**. Initial efforts were focused on setting up fundamental functionalities, authentication flows, and ensuring backend stability. As a result, certain advanced features and integrations had to be deferred to the next sprint.

### Frontend
- End-to-End (E2E) testing is not yet implemented.
- Some parts of Opportunity and Application is still using mock not actual fetched data from backend.
- Dashboard for managing opportunities and applications is not yet built. It lacks the graph for the recent jobs.

### Backend
- Testing for a few special API's are still left
- Some changes in Opportunity model lead to some changes in the testing which are not completed in this sprint.

### Next Steps
The pending features will be prioritized in **Sprint 4**, ensuring:
- Complete E2E testing for all pages

- Full API support for volunteer and organization dashboards.
- Full integration and full coverage for testing in frontend.
- E2E tests for ensuring app stability before deployment.
- Containerization of the application for easy deployment. -->
