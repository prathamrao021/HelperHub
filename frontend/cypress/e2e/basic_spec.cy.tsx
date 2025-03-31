/// <reference types="cypress" />
describe('Initial Load', () => {
    it('should load the homepage and verify key elements', () => {
      cy.visit('/'); 
      
      cy.title().should('include', 'HelperHub'); 
      
      cy.get('nav').should('exist');
      
      cy.contains('Register', { timeout: 5000 }).should('be.visible');
      cy.contains('Login').should('be.visible');
    });
  });