describe('Volunteer Registration', () => {
    it('should register a volunteer successfully', () => {
      // Visit the volunteer registration page
      cy.visit('/register/organization'); // Adjust URL based on your routing setup
  
      // Fill out the registration form
      cy.get('input[name="email"]').type('org@example.com');
      cy.get('input[name="password"]').type('password123');
      cy.get('input[name="name"]').type('My Organization');
      cy.get('input[name="phone"]').type('1234567890');
      cy.get('input[name="location"]').type('123 Main St, City, Country');
      cy.get('textarea[name="description"]').type('We aim to make the world a better place.');
  
    // Submit the form
    cy.get('button[type="submit"]').click();

    // Verify registration was successful and user is redirected to the dashboard
    cy.url().should('include', '/dashboard');
    cy.contains('Dashboard').should('be.visible');
    });
  });