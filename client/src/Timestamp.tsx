import React, { useEffect, useState } from "react";

interface TimestampProps {
  exampleProp: string;
}

export const Timestamp: React.FC<TimestampProps> = ({ }) => {
  const [data, setData] = useState<string>("");

  // use effect
  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await fetch(`/api/timestamp`);
      const data = await response.text();
      setData(data);
    } catch (err: unknown) {
      reportErr(err);
    }
  };

  return <h2>{data}</h2>;
};

const reportErr = (error: unknown) => {
  console.error("An error occurred:", error);
};
