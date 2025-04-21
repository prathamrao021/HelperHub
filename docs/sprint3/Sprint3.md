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

## Sprint 3 Accomplishments

## Frontend Development Accomplishments

### User Interface and Experience
1. **Implemented Comprehensive Application Page**
   - Created a complete volunteer application tracking system with status indicators and filtering
   - Added application withdrawal functionality with confirmation dialogs
   - Integrated loading states and empty state handling for better UX
   - Implemented responsive design for all device types

2. **Built Opportunities Discovery Interface**
   - Developed a fully functional opportunities browsing page for volunteers
   - Implemented advanced filtering by title, location, and date
   - Added card-based UI with clear visual hierarchy
   - Created seamless application submission flow with feedback

3. **Enhanced Authentication System**
   - Refactored auth-context.tsx to support role-based access (Volunteer/Organization)
   - Fixed critical hook usage issues in authentication flow
   - Implemented secure token storage and management
   - Added proper error handling for authentication processes

4. **Implemented Role-Based Protection**
   - Created protected route components with role validation
   - Set up route guards for volunteer and organization specific pages
   - Implemented proper redirection for unauthorized access attempts
   - Added loading states during authentication checks

### Technical Improvements
1. **Refactored Component Architecture**
   - Split complex components into smaller, more maintainable parts
   - Implemented proper state management across components
   - Added proper typing for component props
   - Fixed ref forwarding issues in navigation components

2. **Improved Form Handling**
   - Created reusable form components with validation
   - Implemented file upload functionality with visual feedback
   - Added proper error messaging for form submissions
   - Created multi-step form processes

3. **API Integration**
   - Set up Axios instance with interceptors for API calls
   - Implemented proper error handling for API responses
   - Created token-based authentication for API requests
   - Added loading states during API calls

4. **Enhanced Developer Experience**
   - Added comprehensive type definitions
   - Improved component naming conventions
   - Fixed code style inconsistencies
   - Added proper comments for complex logic

### Frontend Testing

#### Unit Tests

1. **OrganizationRegistration.test.tsx**  
   - `renders the form with fields and submit button`: Ensures all input fields and the submit button are rendered properly.  
   - `validates form fields and shows error message on invalid input`: Tests form validation by entering invalid inputs and checking for error messages.  
   - `validates max length for organization name`: Ensures the organization name does not exceed 100 characters and displays an error if it does.  
   - `submits form and calls registerOrganization on valid input`: Mocks a successful form submission and verifies that `registerOrganization` is called with the correct parameters, and the user is redirected to the dashboard.  

2. **VolunteerRegistration.test.tsx**  
   - `renders the form with fields and submit button`: Verifies that all input fields and the submit button are displayed.  
   - `validates form fields and shows error message on invalid input`: Ensures form validation works by checking for error messages when invalid data is entered.  

#### Cypress E2E Tests  

1. **basic_spec.cy.tsx**  
   - `should load the homepage and verify key elements`: Ensures the homepage loads correctly by verifying the title, navigation bar, and visibility of the "Register" and "Login" buttons.  

2. **login_volunteer_spec.cy.tsx**  
   - `should log in a volunteer and access the dashboard`: Simulates a volunteer logging in and verifies redirection to the dashboard with the expected elements.  

3. **logout_volunteer_spec.cy.tsx**  
   - `should log in as a volunteer, log out, and redirect to the homepage`: Logs in a volunteer, performs logout, and ensures redirection to the homepage with the login button visible.  

4. **register_organization_spec.cy.tsx**  
   - `should register an organization successfully`: Completes the organization registration form and verifies successful redirection to the dashboard.  

5. **register_volunteer_spec.cy.tsx**  
   - `should register a volunteer successfully`: Fills out the volunteer registration form, selects skills, and checks redirection to the dashboard upon successful registration.  


## Backend Development
### Volunteer Management API Endpoints  

### 1. **Volunteer Statistics API**  
**Endpoint**: `GET /volunteers/:volunteer_id/stats`  
**Description**: Retrieve the total number of jobs and hours worked for a volunteer based on accepted applications.  

## Opportunity Management API Endpoints  

### 1. **Last N opportunities for Organization where End Date < Current Date**  
**Endpoint**: `GET /opportunities/organization/:organization_mail/expired`  
**Description**: Retrieve the last 'n' opportunities where the end_date is less than the current date for a specific organization.  

### 2. **Get all Opportunities with number of applications for each**  
**Endpoint**: `GET /opportunities`  
**Description**: Retrieve all opportunities for a specific organization, including the number of applications each opportunity has received. 

### 3. **Last N opportunities for Volunteer where status is Accepted and End Date < Current Date**  
**Endpoint**: `GET /opportunities/volunteer/:volunteer_id/accepted-expired`  
**Description**: Retrieve the last 'n' opportunities for a volunteer where the application was accepted and end_date < current date.  

## Application Management API Endpoints  

### 1. **Create Application**  
**Endpoint**: `POST /applications/`  
**Description**: Submits a new application for an opportunity.  

### 2. **Get All Applications**  
**Endpoint**: `GET /applications/`  
**Description**: Retrieves all applications.  

### 3. **Get Application by ID**  
**Endpoint**: `GET /applications/:id`  
**Description**: Fetches details of a specific application using its ID.  

### 4. **Get Applications by Volunteer ID**  
**Endpoint**: `GET /applications/volunteer/:volunteer_id`  
**Description**: Retrieves all applications submitted by a specific volunteer.  

### 5. **Get Applications by Opportunity ID**  
**Endpoint**: `GET /applications/opportunity/:opportunity_id`  
**Description**: Retrieves all applications for a specific opportunity.  

### 6. **Get Applications by Status**  
**Endpoint**: `GET /applications/status/:status`  
**Description**: Retrieves applications based on their status (e.g., pending, approved, rejected).  

### 7. **Update Application**  
**Endpoint**: `PUT /applications/:id`  
**Description**: Updates the status or details of an application.  

### 8. **Delete Application**  
**Endpoint**: `DELETE /applications/:id`  
**Description**: Deletes an application by its ID.  

### 9. **Get Last N Approved Applications by Volunteer ID**  
**Endpoint**: `GET /applications/volunteer/:volunteer_id/approved`  
**Description**: Retrieves the last N approved applications of a volunteer.  

### Unit Tests  

#### **Volunteer**  
- **TestGetVolunteerStats**  
  - Tests the API for retrieving volunteer statistics (total jobs and hours worked).  
  - Creates test volunteers, opportunities, and applications with different statuses.  
  - Verifies that only "Accepted" applications are counted in stats.  
  - Checks the correct summing of hours worked.  

- **TestGetVolunteerStatsNonExistentVolunteer**  
  - Tests stats retrieval for a non-existent volunteer ID.  
  - Verifies the API returns appropriate zero values rather than errors.  
  - Ensures the endpoint gracefully handles missing data.  

- **TestInvalidUpdateVolunteerData**  
  - Tests updating a volunteer with invalid JSON data.  
  - Verifies proper error handling and status code (400).  

- **TestInvalidLoginRequest**  
  - Tests login endpoint with malformed JSON.  
  - Verifies proper error response and status code.  

#### **Organization**  
- **TestUpdateNonExistentOrganization**  
  - Tests attempting to update an organization that doesn't exist.  
  - Verifies correct error code and message.  
  - Ensures the system properly checks existence before updates.  

- **TestInvalidLoginData**  
  - Tests the login endpoint with invalid JSON format.  
  - Verifies proper error handling and status code (400).  
  - Ensures robust input validation.  

#### **Opportunity**  
- **TestGetLastNExpiredOpportunitiesByOrganization**  
  - Tests the API for retrieving expired opportunities for an organization.  
  - Creates both expired and active test opportunities.  
  - Verifies correct filtering by expiration date.  
  - Checks that the correct number of results is returned.  

- **TestGetOpportunitiesByOrganization**  
  - Tests retrieving opportunities with application counts.  
  - Sets up opportunities with different numbers of applications.  
  - Verifies correct counting of applications per opportunity.  

- **TestGetAvailableOpportunities**  
  - Tests filtering active (non-expired) opportunities.  
  - Verifies that expired opportunities are excluded.  
  - Checks that organization names are included in the response.  

- **TestInvalidRequestGetLastNExpiredOpportunities**  
  - Tests error handling when required parameters are missing or invalid.  
  - Verifies proper error codes and messages.  

- **TestInvalidRequestGetOpportunitiesByOrganization**  
  - Tests error handling when `organization_mail` parameter is missing.  
  - Verifies proper error response.  

- **TestUpdateNonExistentOpportunity**  
  - Tests updating an opportunity that doesn't exist.  
  - Verifies correct 404 response.  

- **TestDeleteNonExistentOpportunity**  
  - Tests deleting a non-existent opportunity.  
  - Verifies proper error handling.  

- **TestCreateOpportunityWithInvalidOrganization**  
  - Tests creating an opportunity with a non-existent organization.  
  - Verifies foreign key constraint handling.  

## Successfully Completed Issues

### Frontend

1. Ingetrage opportunities/available with Opportunity page
2. Implement API Endpoint for Volunteer Opportunities Listing
3. Implement API Endpoint for Organization Opportunities List
4. Change Opportunity model to include start_date, end_date field
5. Fix routes for Manage Project Page
6. Implement Project Detail Page with Application Management
7. Organization Opportunity Management Implementation
8. Role-Based Volunteer Management Platform Implementation
9. Implement Volunteer Applications Page
10. Implement Volunteer Applications Page
11. Implement Volunteer Opportunities Page
12. E2E Test Suites
13. Implementing Test Suites for VolunteerRegistration and OrganizationRegistration

1 - 8 Dhruv Makwana 
9 - 13 Akash Balaji

### Backend

       
1. API endpoint that retrieves available volunteer opportunities for the Opportunities page. This endpoint will be used by volunteers to browse and search for opportunities they can apply to.
2.  API endpoint that retrieves all volunteer opportunities associated with a specific organization. This endpoint is already being called from the frontend Projects page but needs implementation on the backend.
3. The system should provide an API endpoint to fetch the last N opportunities created by an organization where the end_date has already passed (i.e., end_date < current_date). This will help organizations track past opportunities easily.
4. The system should provide an API endpoint to fetch the last N opportunities for which a volunteer's application was accepted and the end_date has already passed (i.e., end_date < current_date). This will help volunteers track their past accepted opportunities.
5. The system should provide an API endpoint to fetch the total number of jobs a volunteer has participated in and the total hours worked across all accepted opportunities. This will help track a volunteer’s contribution over time.
6. Refactor : Current model of opportunity field does not have start_date and end_date field. Add those fields and update opportunity end points.
7. We need to implement Create, Read, Update, and Delete (CRUD) functionality for managing organizations in the system. This will allow users to create new organizations, retrieve organization details, update information, and delete organizations when necessary.
8. Implement Create, Read, Update, and Delete (CRUD) operations for the Application model. This model represents applications submitted by volunteers for various opportunities. It allows volunteers to apply, update their applications, and check their statuses while ensuring organizations can review and manage applications.

9. Implemented Unit tests for -
```
Organization
Volunteer
Application
Category
Opportunity 
```     
Task completed by Pratham: 1-8.
Testing is completed by Nikhil. 
        

## Backlogs that were cleared from sprint 2
### Frontend
1. **E2E Testing**
     - Implemented E2E testing for registration, sign in, sign out
1. **User Dashboard**
     - 

2. **Authentication Integration**
     - Integration of Auth0 for authentication

### Backend
1. **Scheduling APIs**
     - Development of APIs for scheduling functionalities
     - Specialized APIs are pending implementation, including:
     - Fetching all recent jobs completed by a volunteer using `volunteer_id`.
     - Retrieving the total number of jobs and hours worked for a volunteer.
     - Displaying recent opportunities published by organizations.

## Issues Not Completed in Sprint 3

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
- Containerization of the application for easy deployment.
