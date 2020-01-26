import React, { useState } from "react";

export default function LikeButton({ id }) {
  const [likes, setLikes] = useState(() => parseInt(id));
  return (
    <button className="btn btn-secondary" onClick={() => setLikes(draft => setLikes(++draft))}>
      Like ({likes})
    </button>
  );
}
