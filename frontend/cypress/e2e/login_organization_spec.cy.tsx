describe('Volunteer Login Flow', () => {
    it('should log in a volunteer and access the dashboard', () => {
      // Visit the login page
      cy.visit('/login'); // Adjust the URL if necessary based on your routing setup
  
      // Fill out the login form
      cy.get('#email').type('org@example.com');
      cy.get('.space-y-4 > :nth-child(2)').type('password123');
      cy.get('.space-y-4 > .inline-flex').click();
  
      // Verify login was successful and user is redirected to the dashboard
      cy.url().should('include', '/dashboard');
      cy.contains('Dashboard').should('be.visible');
    });
  });