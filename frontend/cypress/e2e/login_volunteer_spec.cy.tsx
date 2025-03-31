describe('Volunteer Login Flow', () => {
    it('should log in a volunteer and access the dashboard', () => {
      cy.visit('/login'); 
  

      cy.get('#email').type('cd@gmail.com');
      cy.get('.space-y-4 > :nth-child(2)').type('12345678');
      cy.get('.space-y-4 > .inline-flex').click();
  
      //redirected to the dashboard
      cy.url().should('include', '/dashboard');
      cy.contains('Dashboard').should('be.visible');
    });
  });