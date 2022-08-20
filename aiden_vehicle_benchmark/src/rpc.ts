declare function GetParentResourceName(): string;

export function rpc<T>(name: string, data: T) {
  fetch(`https://${GetParentResourceName()}/${name}`, {
    body: JSON.stringify(data),
    method: "POST",
    headers: {
      "Content-Type": "application/json; charset=UTF-8",
    },
  });
}
