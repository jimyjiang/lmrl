# Gemini Project Overview

This document provides a brief overview of the `lmrl` project and instructions on how to run it, along with a summary of recent UI improvements.

## Project Description

The `lmrl` project is a Bible search tool with a web-based frontend and a Go backend. It allows users to search for Bible verses or keywords and view the results. It also includes a search history feature.

## How to Run the Project

To set up and run the project, navigate to the project's root directory and execute the following command:

```bash
make dev
```

This command will:
1.  Navigate into the `frontend` directory.
2.  Install frontend dependencies using `npm install`.
3.  Build the frontend application using `vite build`.
4.  Copy necessary frontend assets to the `router/frontend/assets` directory.
5.  Build the Go backend application.
6.  Start the backend server, which will serve the frontend and handle API requests.

The application will typically be accessible at `http://localhost:3001/lmrl/search`.

## Recent UI Improvements

Recent efforts have focused on improving the stability of the user interface, specifically addressing an issue where the left (search history) and right (search results) panels would "jump" or change size unexpectedly when interacting with the search history.

The following changes were implemented in `frontend/src/style.css`, `frontend/src/App.vue`, and `frontend/src/components/BibleSearch.vue` to resolve this:

*   **Global Height Consistency:** `html` and `body` elements are now configured to consistently occupy 100% of the viewport height, providing a stable base for nested height calculations.
*   **Flexbox Layout for Main Application:** The main `#app` container and its direct child `main` are now set up as flex containers with `flex-direction: column`. This ensures that vertical space is distributed predictably among the header and the main content area.
*   **Stable Main Content Area:** The `.google-style-search` container in `BibleSearch.vue` is also a vertical flex container, and its `.main-content` child is configured as a horizontal flex container (`display: flex`, `flex-direction: row`, `flex-grow: 1`, `min-height: 0`). This allows `main-content` to fill the available vertical space and then distribute horizontal space to its children.
*   **Fixed Panel Dimensions and Scrolling:**
    *   The `.history-panel` now has explicit `flex-grow: 0`, `flex-shrink: 0`, and `flex-basis: 200px` (or similar fixed width) to maintain a consistent width.
    *   Both `.history-panel` and `.results-panel` have `height: 100%` and `overflow-y: auto` to ensure they fill the available vertical space within `main-content` and handle their own content scrolling without affecting the overall layout.
    *   `box-sizing: border-box` has been consistently applied to relevant containers for predictable sizing behavior.

These changes collectively aim to create a more stable and predictable layout, preventing unwanted shifts when content changes or scrollbars appear/disappear.