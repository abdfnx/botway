export default function LogoSection() {
  return (
    <section className="flex items-start w-full px-4 mx-auto md:px-0 md:items-center md:w-1/3">
      <div
        className="flex flex-row items-center w-full max-w-sm py-4 mx-auto md:mx-0 my-auto min-w-min relative md:-left-2.5 pt-4 md:py-4 transform origin-left"
        style={{ background: "#13111c" }}
      >
        <div className="flex items-center space-x-1">
          <img
            alt="Botway Logo"
            src="/simple/logo-white.svg"
            className="block w-56"
          />
        </div>
      </div>
    </section>
  );
}
