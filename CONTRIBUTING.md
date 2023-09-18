# Contributing Guide

The following [`opendatahub-io/ai-edge`  GitHub Project board](https://github.com/orgs/opendatahub-io/projects/47) is the source of truth for the work taking place in this repo. You can pick up an issue from the TO DO column and follow the workflow described below. 

<br>

![image](https://github.com/piotrpdev/ai-edge/assets/99439005/dead3499-4005-4dd1-afb5-9a266163f5f4)

<br>

Issues with the `Tracker` label are the high-level longer-term tasks for the PoC, with smaller _“Sub Tasks”_ often listed in the description that should be completed to achieve this goal. These link to other GitHub issues.

<br>

![image](https://github.com/piotrpdev/ai-edge/assets/99439005/459e22f8-2eb4-42c0-8782-851c5540a8ac)

<br>

## Typical Workflow

The easiest way to start contributing is to work on these smaller _“Sub Tasks”_.

<br>

> **NOTE**
> By contributing you agree to the license terms (see [LICENSE](./LICENSE)).

<br>

The general flow of making contributions to the repo goes like this:

- Choose an issue to work on.

- [Assign it to yourself](https://docs.github.com/en/issues/tracking-your-work-with-issues/assigning-issues-and-pull-requests-to-other-github-users).
- If the description is not very detailed, you can improve it yourself. Add a suitable description, link to a User Story, and add acceptance criteria (see next page).
- Add labels and other details e.g. the priority and `kind/documentation` if you will be adding documentation or modifying existing `README` files.
- If there is a roadblock to completing the issue, reach out on the relevant OpenDataHub Slack channels ([invite](https://github.com/opendatahub-io/opendatahub-community)). Someone will gladly try to help. 
  - Sometimes a task or user story ends up not being fully possible to achieve, at least in the way it was intended (like [in this case](https://github.com/opendatahub-io/ai-edge/issues/17)). This is okay, make sure to reach out to others for help instead of trying to do the impossible.
- If your work involves coding (which it probably does) please use the following Git approach:
  - Make a fork of the [`opendatahub-io/ai-edge`](https://github.com/opendatahub-io/ai-edge) repo.
  
  - Create a new branch with a suitable name that describes your work.
  - Create and push commits regularly to save your progress.
  - When you’re ready to make a pull request with your changes, first clean-up by rebasing and squashing the commits. Make sure to use clear and descriptive commit messages.
    - Rebasing and squashing can be tricky so be careful while performing it. You can also learn more about [squashing commits with rebase](https://gitready.com/advanced/2009/02/10/squashing-commits-with-rebase.html).
  - Create the pull request, apply the appropriate labels, and link the relevant GitHub issue to it. Also make sure to include a good description of your changes and screenshots if appropriate.
  - Wait for other members of the team to review your work, you can also tag team members who you think are relevant to your work.
  - Once any conflicts and code suggestions have been resolved, and your work has been approved, you can merge the pull request or wait for somebody else to merge it for you.

<br>

## Examples

<br>

<p align="center"><a target="_blank" href="https://github.com/opendatahub-io/ai-edge/issues/13"><i>Typical “Sub Task” GitHub issue:</i></a></p>

![image](https://github.com/piotrpdev/ai-edge/assets/99439005/cf699d38-9674-47c8-a7e0-6a5a37d0411f)

<br>

<p align="center"><a target="_blank" href="https://github.com/opendatahub-io/ai-edge/pull/75"><i>Typical Pull Request:</i></a></p>

![image](https://github.com/piotrpdev/ai-edge/assets/99439005/3288dede-54f7-4b10-aa06-1f56bd759355)
