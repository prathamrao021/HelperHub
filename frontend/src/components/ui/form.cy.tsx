import React from "react";
import { Form, FormItem, FormLabel, FormControl, FormMessage, FormField } from "./form";
import { Input } from "@/components/ui/input"; 
import { mount } from "cypress/react";
import { useForm } from "react-hook-form";

function TestForm() {
  const form = useForm({
    defaultValues: { name: "" },
    mode: "onSubmit",
  });

  return (
    <Form {...form}>
      <FormField
        name="name"
        control={form.control}
        rules={{ required: "Name is required" }} // Validation rule
        render={({ field }) => (
          <FormItem>
            <FormLabel>Name</FormLabel>
            <FormControl>
              <Input placeholder="Enter name" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )}
      />
      <button type="submit" onClick={form.handleSubmit(() => {})}>Submit</button>
    </Form>
  );
}

describe("Form Component", () => {
  it("renders a form with input fields", () => {
    mount(<TestForm />);
    
    cy.get("[data-slot='form-label']").should("contain.text", "Name");
    cy.get("input").should("have.attr", "placeholder", "Enter name");
  });

  it("should validate required field", () => {
    mount(<TestForm />);
    
    cy.get("button").click(); // Click submit

    cy.get("[data-slot='form-message']")
      .should("exist")
      .should("contain.text", "Name is required");
  });

  it("allows user input", () => {
    mount(<TestForm />);
    
    cy.get("input").type("John Doe").should("have.value", "John Doe");
  });
});
