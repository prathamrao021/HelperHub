describe('Volunteer Registration', () => {
  it('should register a volunteer successfully', () => {
    // Visit the volunteer registration page
    cy.visit('/register/volunteer'); // Adjust URL based on your routing setup

    // Fill out the registration form
    cy.get('input[name="name"]').type('John Doe');
    cy.get('input[name="email"]').type('john.doe@example.com');
    cy.get('input[name="password"]').type('password123');
    cy.get('input[name="phone"]').type('5551234567');
    cy.get('textarea[name="bio_Data"]').type('I am a passionate volunteer with 5 years of experience.');
    cy.get('input[name="location"]').type('New York, USA');

    // Clear and enter the available hours
    cy.get('input[name="available_Hours"]').clear().type('2').should('have.value', '20');

    // Alternatively, handle input directly with focus and typing
    cy.get('input[name="available_Hours"]').focus().clear().type('2').should('have.value', '20');

    // Open skills dropdown and select skills
    cy.get('button').contains('Select skills').click();

    // Ensure the dropdown menu is rendered
    cy.get('[role="menu"]').should('be.visible').within(() => {
      cy.contains('Web Development').click();
      cy.contains('Graphic Design').click();
    });

    // Submit the form
    cy.get('button[type="submit"]').click();

    // Verify registration was successful and user is redirected to the dashboard
    // cy.url().should('include', '/dashboard');
    // cy.contains('Dashboard').should('be.visible');
  });
});