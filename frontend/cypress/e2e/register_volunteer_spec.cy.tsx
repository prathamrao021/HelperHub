describe('Volunteer Registration', () => {
  it('should register a volunteer successfully', () => {
    cy.visit('/register/volunteer'); 

    // registration form
    cy.get('input[name="name"]').type('John Doely');
    cy.get('input[name="email"]').type('jdcb@example.com');
    cy.get('input[name="password"]').type('password1234');
    cy.get('input[name="phone"]').type('5551234568');
    cy.get('textarea[name="bio_Data"]').type('I am a dedicated volunteer with over five years of experience, deeply passionate about making a positive impact in my community through various projects and initiatives.');
    cy.get('input[name="location"]').type('New York, ny, USA');

    cy.get('input[name="available_Hours"]').clear().type('2').should('have.value', '20');


    cy.get('button').contains('Select skills').click();


    cy.get('[role="menu"]').should('be.visible').within(() => {
      cy.contains('Web Development').click();
      cy.contains('Graphic Design').click();
    });

    cy.get('button[type="submit"]').click();

    // Dashboard redirection if succesfull
    cy.url().should('include', '/dashboard');
    cy.contains('Dashboard').should('be.visible');
  });
});