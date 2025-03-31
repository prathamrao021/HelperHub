describe('Volunteer Logout', () => {
    it('should log in as a volunteer, log out, and redirect to the homepage', () => {
      // Visit the login page
      cy.visit('/login');
  
      // Fill out the volunteer login form
      cy.get('#email').type('cd@gmail.com');
      cy.get('.space-y-4 > :nth-child(2)').type('12345678');
      cy.get('.space-y-4 > .inline-flex').click();
  
      // Verify login was successful and user is redirected to the volunteer dashboard
      cy.url().should('include', '/dashboard');
      cy.contains('Dashboard').should('be.visible');
  
      // Log out the volunteer
      cy.get('button').contains('Logout').click(); // Adjust selector based on actual structure
  
      // Verify redirection to the homepage
      cy.url().should('eq', `${Cypress.config().baseUrl}/`);
      cy.contains('Login').should('be.visible'); // Adjust content based on actual homepage
    });
  });