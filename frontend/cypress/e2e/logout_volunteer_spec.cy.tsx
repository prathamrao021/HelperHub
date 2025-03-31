describe('Volunteer Logout', () => {
    it('should log in as a volunteer, log out, and redirect to the homepage', () => {
      cy.visit('/login');

      cy.get('#email').type('cd@gmail.com');
      cy.get('.space-y-4 > :nth-child(2)').type('12345678');
      cy.get('.space-y-4 > .inline-flex').click();
  
      //volunteer dashboard
      cy.url().should('include', '/dashboard');
      cy.contains('Dashboard').should('be.visible');
  
      // Log out the volunteer
      cy.get('button').contains('Logout').click(); 
  
      cy.url().should('eq', `${Cypress.config().baseUrl}/`);
      cy.contains('Login').should('be.visible'); // Verify that the login button is visible again
    });
  });