let standard_input = process.stdin;
standard_input.setEncoding("utf-8");
standard_input.on("data", function(data) {
  data = data.split("\n");

  let inp = data[1].split(" ").map(x => parseInt(x));
  let s = parseInt(data[0].split(" ")[1]);
  let out = 0;

  let [left, right] = halve(inp);
  let lMap = ss(left, s);
  let rMap = ss(right, s);

  for (const [lk, lv] of Object.entries(lMap)) {
    if (s - parseInt(lk) in rMap) {
      out += lv * rMap[s - parseInt(lk)];
    }
  }
  console.log(out);
  process.exit();
});

let ss = (inp, s, idx = 0, sum = 0, acc = { 0: 1 }) => {
  if (sum > s) {
    return;
  }

  if (idx === inp.length) {
    return acc;
  }

  if (sum + inp[idx] <= s) {
    acc = add2Map(sum + inp[idx], acc);
    ss(inp, s, idx + 1, sum + inp[idx], acc);
  }

  ss(inp, s, idx + 1, sum, acc);
  return acc;
};

let add2Map = (sum, sums) => {
  if (sum in sums) {
    sums[sum]++;
  } else {
    sums[sum] = 1;
  }
  return sums;
};

let halve = l => {
  let mid = Math.floor(l.length / 2);
  return [l.slice(0, mid), l.slice(mid, l.length)];
};
