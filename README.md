# fivem-projects

I've been learning about [FiveM](https://fivem.net/) recently. This repo is just a monorepo/scratchpad of all the random resources I'm building! I originally moved to a monorepo so that I can stop spamming my GitHub profile with so many tiny repos.

Most of these have been streamed live on my [Twitch](https://twitch.tv/Aiden), if these interest you, stop by!

## Installing

If you wanted to use one of these resources, I'd recommend adding it to your `server-data` by creating this monorepo as a [resource category](https://docs.fivem.net/docs/scripting-manual/introduction/introduction-to-resources/#resource-directories), that way you can pull in any updates far easier than cloning/moving files out of the repo.

You could do this with the following git command:

```bash
git clone git@github.com:aidenwallis/fivem-projects.git [aiden_fivem_projects]
```

This would clone to a directory called `[aiden_fivem_projects]`. You can then edit your `server.cfg` and `ensure` any module you want to use.

**Some modules will require you building them, refer to each projects' README for detailed instructions.**

## Structure

You should treat each project directory as its own separate repo - open it in a separate editor per project directory, etc.
