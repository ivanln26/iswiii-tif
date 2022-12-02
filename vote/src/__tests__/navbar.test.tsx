import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import Navbar from "@/components/navbar";

describe("Navbar", () => {
  it("renders navbar", () => {
    render(<Navbar />);

    const homeLink: HTMLAnchorElement = screen.getByText("Home");
    const listLink: HTMLAnchorElement = screen.getByText("List");
    const percentagesLink: HTMLAnchorElement = screen.getByText("Percentages");

    expect(homeLink).toBeInTheDocument();
    expect(listLink).toBeInTheDocument();
    expect(percentagesLink).toBeInTheDocument();

    expect(homeLink.href).toMatch(/\/$/);
    expect(listLink.href).toMatch(/\/list$/);
    expect(percentagesLink.href).toMatch(/\/percentages$/);
  });
});
