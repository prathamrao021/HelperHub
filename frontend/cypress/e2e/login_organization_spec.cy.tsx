describe('Volunteer Login Flow', () => {
    it('should log in a volunteer and access the dashboard', () => {
      cy.visit('/login'); 
  
      cy.get('#email').type('org@example.com');
      cy.get('.space-y-4 > :nth-child(2)').type('password123');
      cy.get('.space-y-4 > .inline-flex').click();
  
      cy.url().should('include', '/dashboard');
      cy.contains('Dashboard').should('be.visible');
    });
  });