import { useState } from "react";

function App() {
  const URL_API = "http://localhost:8080";

  const [data, setData] = useState(null);
  const [idUser, setIdUSer] = useState(1);
  const [isOpenCard, setIsOpenCard] = useState(false);
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [typeCard, setTypeCard] = useState("");

  const getRandomCard = async (idUser) => {
    try {
      const url = `${URL_API}/random-card/${idUser}`;
      const response = await fetch(url);
      if (!response.ok) throw new Error('Error fetching data');
      
      const dataResponse = await response.json();
      console.log(dataResponse)
      setData(dataResponse);

      if (dataResponse) {
        setTitle(dataResponse.title);
        setContent(dataResponse.content);
        setTypeCard(dataResponse.type.name);
        setIsOpenCard(true);
      }
    } catch (error) {
      console.error("Error fetching random card:", error);
    }
  };

  return (
    <div className="min-h-screen bg-slate-800">
      <div className="mx-auto my-auto max-w-xl p-3">
        <h1 className="mb-4 pt-8 text-3xl text-center font-extrabold text-gray-900 dark:text-white md:text-5xl lg:text-6xl hover:scale-105 transition">
          <span className="text-transparent bg-clip-text bg-gradient-to-r to-emerald-600 from-sky-400">
            Interesting Conversations
          </span>
        </h1>
        <p className="text-lg text-center font-normal text-gray-500 lg:text-xl dark:text-gray-400 py-2">
          Connect Deeper: Discover, Share, Engage with Cards!
        </p>
      </div>

      <div>
        <button
          type="button"
          className="block mx-auto text-gray-900 bg-gradient-to-r from-red-200 via-red-300 to-yellow-200 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-red-100 dark:focus:ring-red-400 font-medium rounded-lg text-sm px-5 py-2.5 text-center mb-2 hover:scale-110 transition"
          onClick={() => getRandomCard(idUser)}
        >
          Random Topic
        </button>
      </div>

      {isOpenCard && (
        <div className="max-w-sm mx-auto my-8 p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
          {typeCard === "question" && (
            <svg
              className="w-7 h-7 text-gray-500 dark:text-gray-400 mb-3"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              fill="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10-4.477 10-10 10S2 17.523 2 12Zm9.008-3.018a1.502 1.502 0 0 1 2.522 1.159v.024a1.44 1.44 0 0 1-1.493 1.418 1 1 0 0 0-1.037.999V14a1 1 0 1 0 2 0v-.539a3.44 3.44 0 0 0 2.529-3.256 3.502 3.502 0 0 0-7-.255 1 1 0 0 0 2 .076c.014-.398.187-.774.48-1.044Zm.982 7.026a1 1 0 1 0 0 2H12a1 1 0 1 0 0-2h-.01Z"
              />
            </svg>
          )}
          {typeCard === "topic" && (
            <svg
              className="w-7 h-7 text-gray-500 dark:text-gray-400 mb-3"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              fill="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                d="M11 4.717c-2.286-.58-4.16-.756-7.045-.71A1.99 1.99 0 0 0 2 6v11c0 1.133.934 2.022 2.044 2.007 2.759-.038 4.5.16 6.956.791V4.717Zm2 15.081c2.456-.631 4.198-.829 6.956-.791A2.013 2.013 0 0 0 22 16.999V6a1.99 1.99 0 0 0-1.955-1.993c-2.885-.046-4.76.13-7.045.71v15.081Z"
              />
            </svg>
          )}
          {typeCard === "dynamic" && (
            <svg
              className="w-7 h-7 text-gray-500 dark:text-gray-400 mb-3"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              fill="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                d="M3.559 4.544c.355-.35.834-.544 1.33-.544H19.11c.496 0 .975.194 1.33.544.356.35.559.829.559 1.331v9.25c0 .502-.203.981-.559 1.331-.355.35-.834.544-1.33.544H15.5l-2.7 3.6a1 1 0 0 1-1.6 0L8.5 17H4.889c-.496 0-.975-.194-1.33-.544A1.868 1.868 0 0 1 3 15.125v-9.25c0-.502.203-.981.559-1.331ZM7.556 7.5a1 1 0 1 0 0 2h8a1 1 0 0 0 0-2h-8Zm0 3.5a1 1 0 1 0 0 2H12a1 1 0 1 0 0-2H7.556Z"
              />
            </svg>
          )}

          <h5 className="mb-2 text-2xl font-semibold tracking-tight text-gray-900 dark:text-white">
            {title}
          </h5>
          <p className="mb-3 font-normal text-gray-500 dark:text-gray-400">
            {content}
          </p>
          <a
            href="https://github.com/ManuEduardo/random-topic"
            className="inline-flex font-medium items-center text-blue-600 hover:underline"
          >
            Github
            <svg
              className="w-3 h-3 ms-2.5 rtl:rotate-[270deg]"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 18 18"
            >
              <path
                stroke="currentColor"
                d="M15 11v4.833A1.166 1.166 0 0 1 13.833 17H2.167A1.167 1.167 0 0 1 1 15.833V4.167A1.166 1.166 0 0 1 2.167 3h4.618m4.447-2H17v5.768M9.111 8.889l7.778-7.778"
              />
            </svg>
          </a>
        </div>
      )}
    </div>
  );
}

export default App;
