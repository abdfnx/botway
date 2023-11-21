export const LogoSection = () => {
  return (
    <section className="flex items-start w-full place-content-center px-4 mx-auto md:px-0 md:items-center md:w-1/3">
      <div
        style={{ background: "#13111c" }}
        className="flex flex-row place-content-center items-center w-full max-w-sm py-4 mx-auto md:mx-0 my-auto min-w-min relative pt-4 md:py-4 transform md:origin-right"
      >
        <div className="flex items-center place-content-center space-x-1">
          <img
            alt="Botway Logo"
            src="/simple/logo-white.svg"
            className="block w-72"
          />
        </div>
      </div>
    </section>
  );
};
