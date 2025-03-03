import React from "react";
import { Button } from "./button";
import { mount } from "cypress/react";

describe("Button Component", () => {
  it("renders correctly", () => {
    mount(<Button>Click Me</Button>);
    cy.get("[data-slot='button']").should("contain.text", "Click Me");
  });

  it("should apply default variant and size", () => {
    mount(<Button>Click Me</Button>);
    cy.get("[data-slot='button']").should("have.class", "bg-primary");
    cy.get("[data-slot='button']").should("have.class", "h-9");
  });

  it("should apply custom variant and size", () => {
    mount(<Button variant="destructive" size="lg">Delete</Button>);
    cy.get("[data-slot='button']").should("have.class", "bg-destructive");
    cy.get("[data-slot='button']").should("have.class", "h-10");
    cy.get("[data-slot='button']").should("contain.text", "Delete");
  });

  it("should be disabled when the disabled prop is passed", () => {
    mount(<Button disabled>Disabled</Button>);
    cy.get("[data-slot='button']").should("be.disabled");
  });

  it("should trigger onClick when clicked", () => {
    const handleClick = cy.spy();
    mount(<Button onClick={handleClick}>Click Me</Button>);

    cy.get("[data-slot='button']").click();
    cy.wrap(handleClick).should("have.been.calledOnce");
  });
});
