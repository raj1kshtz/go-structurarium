# GitHub Repository Setup Guide

This guide will help you configure your repository for automated testing, coverage reporting, and branch protection.

## 📋 Table of Contents

1. [Configuring Branch Protection](#configuring-branch-protection)
2. [Setting Up Code Coverage](#setting-up-code-coverage)
3. [Testing the Setup](#testing-the-setup)
4. [GitHub Notifications](#github-notifications)

## 🔒 Configuring Branch Protection

### Step 1: Navigate to Branch Settings

1. Go to **Settings** → **Branches**
2. Click **Add rule** under "Branch protection rules"

### Step 2: Configure Protection Rules

**Branch name pattern:** `main`

Enable the following:

```
✅ Require a pull request before merging
   ├─ Require approvals: 1
   ├─ Dismiss stale pull request approvals when new commits are pushed
   └─ Require approval of the most recent reviewable push

✅ Require status checks to pass before merging
   ├─ Require branches to be up to date before merging
   └─ Status checks that are required:
      - test (1.21)
      - test (1.22)
      - test (1.23)
      - lint
      - build
      - test-summary

✅ Require conversation resolution before merging

✅ Require linear history (optional but recommended)

✅ Do not allow bypassing the above settings
   (Recommended: This applies rules to administrators too)

✅ Restrict who can push to matching branches (optional)

✅ Block force pushes

❌ Allow deletions (leave unchecked)

❌ Allow force pushes (leave unchecked)
```

### Step 3: Save Changes

Click **Create** or **Save changes**

## 📊 Setting Up Code Coverage

### Option 1: Using Codecov (Recommended)

1. Go to https://codecov.io/
2. Sign in with GitHub
3. Add your repository: `go-structurarium`
4. Copy the upload token (optional, but recommended)
5. Add it as a GitHub secret:
   - **Name:** `CODECOV_TOKEN`
   - **Value:** Your Codecov token

The workflow is already configured to upload coverage automatically.

### Option 2: GitHub Actions Summary Only

No additional setup needed! Coverage reports will appear in:
- PR comments
- GitHub Actions summary page
- Workflow logs

## 🧪 Testing the Setup

### Test 1: Verify Workflows

1. Push your changes:
   ```bash
   git add .
   git commit -m "Setup GitHub Actions and workflows"
   git push origin main
   ```

2. Go to **Actions** tab on GitHub
3. You should see workflows running

### Test 2: Test Pull Request Workflow

1. Create a test branch:
   ```bash
   git checkout -b test/pr-workflow
   echo "# Test" >> test.txt
   git add test.txt
   git commit -m "Test PR workflow"
   git push origin test/pr-workflow
   ```

2. Go to GitHub and create a Pull Request
3. Check that:
   - ✅ Tests run automatically
   - ✅ Coverage report appears as a comment
   - ✅ All status checks show up
   - ✅ You need approval to merge

### Test 3: Verify Branch Protection

1. Try to push directly to main:
   ```bash
   git checkout main
   echo "test" >> README.md
   git commit -am "Direct push test"
   git push origin main
   ```

2. You should see:
   ```
   remote: error: GH006: Protected branch update failed
   ```

3. This confirms branch protection is working! ✅

## 🎯 What You'll Get

### On Every Pull Request:
- ✅ Automated test runs on Go 1.21, 1.22, and 1.23
- ✅ Linting checks
- ✅ Build verification
- ✅ Coverage report posted as PR comment
- ✅ Coverage badge in summary
- ✅ All checks must pass before merge

### On Issues:
- ✅ GitHub notification system alerts you of new issues
- ✅ Email notifications via GitHub (configure in your GitHub settings)

## 🔔 GitHub Notifications

You can receive notifications for issues and PRs through GitHub's built-in notification system:

1. Go to **Settings** (your profile settings, not repository)
2. Click **Notifications** in the left sidebar
3. Configure how you want to receive notifications:
   - Web notifications
   - Email notifications
   - Mobile notifications (via GitHub app)

### Watch Your Repository

To get notified about all activity:
1. Go to your repository
2. Click **Watch** at the top
3. Choose notification level:
   - **All Activity** - Get notified of all issues, PRs, and discussions
   - **Participating and @mentions** - Only when you're involved
   - **Ignore** - No notifications

## 🔧 Troubleshooting

### Status Checks Not Appearing

**Fix:**
1. Wait for the first workflow run to complete
2. Status checks appear after first successful run
3. Then add them to branch protection rules

### Cannot Merge PR (Even with Approval)

**Possible causes:**
1. Tests haven't completed
2. Tests failed
3. Conversations not resolved
4. Branch not up to date with main

**Fix:**
1. Wait for all checks to complete
2. Fix any test failures
3. Resolve all review comments
4. Update branch: `git merge main` or use "Update branch" button

## 📚 Additional Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Branch Protection Rules](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches/about-protected-branches)
- [Codecov Documentation](https://docs.codecov.com/docs)

## 💡 Tips

1. **Review workflow runs regularly** to ensure everything is working
2. **Update Go versions** in workflows when new versions are released
3. **Monitor coverage trends** to maintain code quality
4. **Respond to issues promptly** to build community trust

## 🎉 You're All Set!

Your repository now has:
- ✅ Automated testing on every PR
- ✅ Coverage reporting with detailed summaries
- ✅ Branch protection preventing direct pushes to main
- ✅ Required code reviews
- ✅ Professional issue templates
- ✅ Linting and build verification

Happy coding! 🚀

