describe('Volunteer Registration', () => {
    it('should register a volunteer successfully', () => {
      cy.visit('/register/organization'); 
  
      // Fill out the registration form
      cy.get('input[name="email"]').type('org56@example.com');
      cy.get('input[name="password"]').type('password1234');
      cy.get('input[name="name"]').type('My Organizations');
      cy.get('input[name="phone"]').type('1234567890');
      cy.get('input[name="location"]').type('123 Main St, Ciity, Country');
      cy.get('textarea[name="description"]').type('We aim to make the world a better place by.');
  

    cy.get('button[type="submit"]').click();

    // dashboard if success
    cy.url().should('include', '/dashboard');
    cy.contains('Dashboard').should('be.visible');
    });
  });