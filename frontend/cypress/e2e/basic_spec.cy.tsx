describe('Initial Load', () => {
    it('should load the homepage and verify key elements', () => {
      // Visit the homepage
      cy.visit('/'); 
      
      // Check that the page loaded successfully
      cy.title().should('include', 'HelperHub'); 
      
      // Verify navigation elements exist
      cy.get('nav').should('exist');
      
      // Verify important UI elements
      cy.contains('Register', { timeout: 5000 }).should('be.visible');
      cy.contains('Login').should('be.visible');
    });
  });