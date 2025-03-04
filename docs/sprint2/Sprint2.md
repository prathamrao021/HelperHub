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

## Sprint 2 Accomplishments

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

6. **Unit Tests**
    - Implemented Unit test for -
        Applications -

            TestCreateApplication: Validates that volunteers can successfully apply for opportunities with proper validation and persistence.

            TestGetAllApplications: Ensures the system can retrieve all applications across the platform.
        
            TestGetApplicationByID: Verifies individual applications can be accessed by their unique ID with all relevant data.
        
            TestGetApplicationsByVolunteerID: Tests filtering applications by volunteer, enabling volunteers to view their application history.
        
            TestGetApplicationsByOpportunityID: Validates filtering applications by opportunity, allowing organizations to review all applicants.
        
            TestGetApplicationsByStatus: Confirms applications can be filtered by status (pending, accepted, rejected), supporting workflow management.
        
            TestUpdateApplication: Ensures applications can be updated, particularly for status changes during the review process.
        
            TestDeleteApplication: Verifies applications can be removed from the system when necessary.
        
            TestGetNonExistentApplication: Tests proper error handling for non-existent application lookups.
        
            TestInvalidApplicationData: Validates error handling for malformed application submissions.
      
      Category -

          TestCreateCategory: Ensures the system can initialize with a predefined set of volunteering categories.

          TestGetCategories: Verifies that all categories can be retrieved for display in the UI.

          TestCreateCategoryIdempotence: Confirms that the category creation process is idempotent, preventing duplicate categories when the initialization process runs multiple times.

          TestCategoryDatabaseError: Validates proper error handling when database operations fail.
      
      Opportunity -

          TestCreateOpportunity: Ensures organizations can create new volunteer opportunities with proper validation and persistence.

          TestGetOpportunity: Verifies that opportunities can be retrieved by ID with all data correctly populated.

          TestUpdateOpportunity: Tests the ability to modify opportunity details such as title, description, and requirements.

          TestDeleteOpportunity: Confirms opportunities can be removed from the system when they're no longer available.

          TestGetNonExistentOpportunity: Tests proper error handling when attempting to access non-existent opportunities.

          TestInvalidOpportunityData: Validates error handling for malformed opportunity submissions.
      
      Organization -

          TestCreateOrganization: Ensures organizations can be successfully created with proper data validation, password hashing, and persistence.

          TestGetOrganization: Verifies that organizations can be retrieved by email address with all fields correctly populated.

          TestUpdateOrganization: Tests partial updates to organization details while maintaining unchanged fields.

          TestUpdateOrganizationWithPassword: Specifically tests password updates, ensuring passwords are properly hashed and not stored in plaintext.

          TestDeleteOrganization: Confirms organizations can be deleted from the system.

          TestLoginOrganization: Validates the authentication flow with correct credentials returns proper user data and authentication tokens.

          TestLoginOrganizationInvalidPassword: Verifies authentication fails with incorrect password.

          TestLoginOrganizationInvalidEmail: Ensures authentication fails for non-existent organizations.

          TestGetNonExistentOrganization: Validates proper error handling when attempting to retrieve organizations that don't exist.

          TestInvalidOrganizationData: Tests error handling when invalid data is submitted.
      
       Volunteer -

          User Creation (TestCreateVolunteer)

            Tests creation of new volunteers with proper password hashing
            Verifies the volunteer is saved in the database with all correct fields


          User Retrieval (TestGetVolunteer)

            Tests fetching a volunteer by email
            Verifies all volunteer data is returned correctly, including complex fields like category lists
            

          User Update (TestUpdateVolunteer)

            Tests updating volunteer profiles
            Verifies changes are properly saved to the database


          Password Management (TestUpdateVolunteerWithPassword)

            Tests updating a volunteer's password
            Verifies passwords are properly hashed and stored securely


          User Deletion (TestDeleteVolunteer)

            Tests removing volunteers from the system
            Verifies the volunteer is completely removed from the database


          Authentication (TestLoginVolunteer)

            Tests volunteer login with correct credentials
            Verifies correct user data is returned upon successful authentication


          Authentication Security (TestLoginVolunteerInvalidPassword, TestLoginVolunteerInvalidEmail)

            Tests login failure with incorrect password
            Tests login failure with non-existent email


          Error Handling (TestGetNonExistentVolunteer, TestInvalidVolunteerData)

            Tests API responses for non-existent volunteers
            Tests system's handling of invalid JSON data
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

// List unit tests and Cypress test for frontend
<!-- 1. **Technical Setup**
    - Configured TypeScript with Material UI
    - Established project structure
    - Set up linting and formatting

2. **UI Components**
    - Implemented basic routing
    - Created user login page
    - Developed registration page with validation
    - Built navigation bar with search and theme toggle
    - Designed basic profile page layout

3. **User Interface Features**
    - Homepage layout optimization
    - Profile management pages
    - Navigation system implementation -->

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





<!-- 1. **Swagger documentation integration** - Pratham
2. **Project initialization files** - Pratham
3. **User API endpoints** - Nikhil & Pratham
     - User Create - Pratham
     - User Update - Nikhil
     - User Delete - Pratham
     - User Validate - Nikhil
4. **Database connection** - Pratham -->

<!-- ## Issues Not Completed 

### Frontend
<!-- 1. **User Dashboard**
     - Full dashboard with functionality to update personal information

2. **Authentication Integration**
     - Integration of Auth0 for authentication

### Backend
1. **Scheduling APIs**
     - Development of APIs for scheduling functionalities -->
