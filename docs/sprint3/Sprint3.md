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

## Backend Development
## Volunteer Management API Endpoints  

### 1. **Create Volunteer**  
**Endpoint**: `POST /volunteers/create`  
**Description**: Registers a new volunteer with required details.  

### 2. **Delete Volunteer**  
**Endpoint**: `DELETE /volunteers/delete/:volunteer_mail`  
**Description**: Deletes a volunteer account based on the provided email.  

### 3. **Update Volunteer Details**  
**Endpoint**: `PUT /volunteers/update/:volunteer_mail`  
**Description**: Updates a volunteer’s profile information.  

### 4. **Get Volunteer Details**  
**Endpoint**: `GET /volunteers/get/:volunteer_mail`  
**Description**: Retrieves details of a specific volunteer by email.  

### 5. **Volunteer Login**  
**Endpoint**: `POST /login/volunteer`  
**Description**: Authenticates a volunteer using email and password.  


## Organization Management API Endpoints  

### 1. **Create Organization**  
**Endpoint**: `POST /organizations/create`  
**Description**: Registers a new organization with required details.  

### 2. **Delete Organization**  
**Endpoint**: `DELETE /organizations/delete/:organization_mail`  
**Description**: Deletes an organization account based on the provided email.  

### 3. **Update Organization Details**  
**Endpoint**: `PUT /organizations/update/:organization_mail`  
**Description**: Updates an organization’s profile information.  

### 4. **Get Organization Details**  
**Endpoint**: `GET /organizations/get/:organization_mail`  
**Description**: Retrieves details of a specific organization by email.  

### 5. **Organization Login**  
**Endpoint**: `POST /login/organization`  
**Description**: Authenticates an organization using email and password.  

## Opportunity Management API Endpoints  

### 1. **Create Opportunity**  
**Endpoint**: `POST /opportunities/create`  
**Description**: Creates a new volunteering opportunity.  

### 2. **Delete Opportunity**  
**Endpoint**: `DELETE /opportunities/delete/:id`  
**Description**: Deletes an opportunity based on the provided ID.  

### 3. **Update Opportunity Details**  
**Endpoint**: `PUT /opportunities/update/:id`  
**Description**: Updates details of an existing opportunity.  

### 4. **Get Opportunity Details**  
**Endpoint**: `GET /opportunities/get/:id`  
**Description**: Retrieves details of a specific opportunity by ID.  

## Category Management API Endpoints  

### 1. **Create Category**  
**Endpoint**: `POST /categories/create`  
**Description**: Creates a new category for opportunities.  

### 2. **Get All Categories**  
**Endpoint**: `GET /categories/get`  
**Description**: Retrieves all available categories.  

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

### Successfully Completed Issues For Backend

1. Implemented Model for -
   
       Organization
       Volunteer
       Application
       Category
       Opportunity 

3. Implemented Unit tests for -
   
       Organization
       Volunteer
       Application
       Category
       Opportunity 
        

### Frontend Development

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

## Successfully Completed Issues

### Frontend

1. Configured the project environment and dependencies to support the migration from Material UI to Shadcn, so that the team can seamlessly implement Shadcn components while maintaining code consistency, theming, and performance.
   
2. Implemented Theme Switcher in Navbar for Light/Dark/System Mode Toggles.
   
3. Refactored the Navbar component to use Shadcn instead of Material UI, so that the navigation aligns with the new design system, improves visual consistency, and enhances user experience.
   
4. Rebuilt the Hero section on the landing page using Shadcn components, so that it aligns with the modern design system, improves visual appeal, and ensures consistency with other migrated components.
   
5. Rebuilt the Register page using Shadcn components and Tailwind CSS, so that the registration flow aligns with the application’s modern design system, improves accessibility, and ensures a seamless user experience alongside the migrated Login page.

6. Fixed all the Navigation Links and Button Functionality in Hero Section and Navigation Bar for a seamless experience for the user.
    
7. Implemented Volunteer Registration Page with Profile Details and Skill Selection so that the organization can review my profile and match me with relevant opportunities.
    
8. Implement Organization Registration Page with Essential Details and Logo Upload so that my organization can be onboarded to the platform and access relevant features.

9. Migrated existing core components (Login, Register, Landing Page, Hero Page) from Material UI to Shadcn, so that the application achieves a more elegant, modern, and cohesive design system while improving maintainability and alignment with current design trends.
    
10. Developed a unified, secure, and visually consistent experience across registration, authentication, navigation, and dashboards,
so that user can efficiently interact with the platform based on my role, access tailored features, and enjoy a modern, intuitive interface.

11. Implement Authentication Context for Role-Based Access Control and Protected Routes so that the application can handle Organization/Volunteer logins, protect routes based on roles, and maintain consistent user sessions.
    
12. Integrated a component that allows for the following functionality:
    
        Editing details in organization profile
        Saving the changes made and updating the same in the Database
    
13. Add Edit Organization Dialog in Organization Profile so that the Dashboard automatically updates the displayed information on dashboard.

14. Refactored Navbar component to use consistent button naming and improve skill selection UI in EditVolunteerProfile
    
15. Implemented Delete Button for users to delete their account.

16. Refactored the NavBar component to be more flexible to accomodate for seamless login and logout.

Tasks done by Dhruv : 1-5, 9, 14-16
Tasks done by Akash : 6-8, 10-13

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

## Issues Not Completed in Sprint 2

### Reason for Incompletion
The pending tasks from Sprint 2 were not completed due to **time constraints and prioritization of core features**. Initial efforts were focused on setting up fundamental functionalities, authentication flows, and ensuring backend stability. As a result, certain advanced features and integrations had to be deferred to the next sprint.

### Frontend
- End-to-End (E2E) testing is not yet implemented.
- Forms for creating an opportunity are incomplete.
- Functionality to apply to an opportunity is missing.
- Dashboard for managing opportunities and applications is not yet built.

### Backend
- Specialized APIs are pending implementation, including:
  - Fetching all recent jobs completed by a volunteer using `volunteer_id`.
  - Retrieving the total number of jobs and hours worked for a volunteer.
  - Displaying recent opportunities published by organizations.

### Next Steps
The pending features will be prioritized in **Sprint 4**, ensuring:
- Complete E2E testing for all pages

- Full API support for volunteer and organization dashboards.
- Completion of opportunity-related forms and workflows.
- E2E tests for ensuring app stability before deployment.
